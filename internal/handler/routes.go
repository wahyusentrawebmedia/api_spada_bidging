package handler

import (
	"api/spada/internal/database"
	"api/spada/internal/middleware"
	"api/spada/internal/repository"
	"api/spada/internal/service"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	repoKategori := repository.NewPostgresConfigRepository(database.DB)
	servicesConfig := service.NewPostgresConfigService(*repoKategori)

	postgresConfigHandler := NewPostgresConfigHandler(*servicesConfig)
	userHandler := NewUserHandler(service.NewUserService())

	fakultasHandler := NewFakultasHandler(*service.NewMoodleFakultasService())
	prodiHandler := NewProdisHandler(*service.NewMoodleProdiService())
	tahunHandler := NewTahunAkademikHandler(*service.NewMoodleTahunAkademikService())
	semesterHandler := NewSemesterHandler(*service.NewMoodleSemesterService())
	makulHandler := NewMoodleMakulHandler(service.NewMoodleMakulService())
	categoriHandler := NewCategoriesHandler(*service.NewMoodleCategoriesService())
	groupHandler := NewGroupsHandler(*service.NewMoodleGroupsService())

	{
		appSecureuser := app.Group("/dosen", middleware.JWTCheckMiddlewareUser())

		// User CRUD
		appSecureuser.Get("/", userHandler.GetDetail)
		appSecureuser.Post("/", userHandler.UpdateSingle)
	}

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "pong"})
	})

	{
		appAkademik := app.Group("/akademik", middleware.JWTCheckMiddleware())

		// User CRUD
		appAkademik.Get("/users", userHandler.GetAllUsers)
		appAkademik.Post("/users", userHandler.UpdateSingle)

		// Fakultas
		fakultasRoute := appAkademik.Group("/fakultas")

		fakultasRoute.Get("/", fakultasHandler.GetFakultas)
		fakultasRoute.Post("/", fakultasHandler.CreateFakultas)
		fakultasRoute.Post("/sync", fakultasHandler.SyncFakultas)

		// Prodi
		prodiRoute := appAkademik.Group("/fakultas/:id/")

		prodiRoute.Get("/prodi", prodiHandler.GetProdis)
		prodiRoute.Post("/prodi", prodiHandler.CreateProdis)
		prodiRoute.Post("/prodi-sync", prodiHandler.SyncProdis)

		// Tahun Akademik
		tahunAkademikRoute := appAkademik.Group("/fakultas/:id/prodi/:prodi_id/tahun")

		tahunAkademikRoute.Get("/", tahunHandler.CreatedTahunAkademik)
		tahunAkademikRoute.Post("/", tahunHandler.CreatedTahunAkademik)
		tahunAkademikRoute.Post("/tahun/sync", tahunHandler.SyncdTahunAkademik)

		// Semester
		semesterRoute := appAkademik.Group("/fakultas/:id/prodi/:prodi_id/tahun/:tahun_id/semester")

		semesterRoute.Get("/", semesterHandler.GetdSemester)
		semesterRoute.Post("/", semesterHandler.CreatedSemester)
		semesterRoute.Post("/sync", semesterHandler.SyncdSemester)
		semesterRoute.Get("/:semester_id", semesterHandler.GetDetailSemester)

		// Makul
		makulRoute := appAkademik.Group("/fakultas/:id/prodi/:prodi_id/tahun/:tahun_id/semester/:semester_id/makul")

		// makulRoute.Get("/", makulHandler.GetMakul)
		// makulRoute.Post("/", makulHandler.CreateMakul)
		makulRoute.Post("/sync", makulHandler.SyncMakul)

		// Categories
		categoryRoute := appAkademik.Group("/categories")

		categoryRoute.Get("/prefix/:prefix", categoriHandler.GetCategoriesWithPrefix)
		categoryRoute.Get("/", categoriHandler.GetCategories)
		categoryRoute.Post("/", categoriHandler.CreateCategories)
		categoryRoute.Post("/sync", categoriHandler.SyncCategories)

		// Makul per Semester
		makulCategoriesRoute := categoryRoute.Group("/:semester_id/makul")
		makulCategoriesRoute.Post("/sync", makulHandler.SyncMakul)

		// Groups
		groupsRoute := appAkademik.Group("/groups")

		groupsRoute.Get("/categories/:categories_id", groupHandler.GetGroupsByCategoriesID)

		// Dosen
		DosenRoute := appAkademik.Group("")

		DosenRoute.Get("/dosen", userHandler.GetAllUsers)
		DosenRoute.Post("/dosen-sync", userHandler.Sync)
	}

	{

		appSecure := app.Group("config", middleware.JWTCheckMiddlewareUser())

		// Postgres Config CRUD
		appSecure.Get("/config/postgres", postgresConfigHandler.ListPostgresConfigs)
		appSecure.Post("/config/postgres", postgresConfigHandler.CreatePostgresConfig)
		appSecure.Get("/config/postgres/:id", postgresConfigHandler.GetPostgresConfig)
		appSecure.Put("/config/postgres/:id", postgresConfigHandler.UpdatePostgresConfig)
		appSecure.Delete("/config/postgres/:id", postgresConfigHandler.DeletePostgresConfig)
	}

}
