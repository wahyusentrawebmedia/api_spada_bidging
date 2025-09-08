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
	// moodleHandler := NewMoodleHandler(*service.NewMoodleService())

	appSecure := app.Use(middleware.JWTCheckMiddleware())

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "pong"})
	})

	appAkademik := appSecure.Group("/akademik")

	{
		// User CRUD
		appAkademik.Get("/users", userHandler.GetAllUsers)
		appAkademik.Post("/users", userHandler.Index)
		// appAkademik.Post("/user/sync", userHandler.Sync)
		// appSecure.Post("/user/reset", userHandler.Reset)

		// Moodle: Update password user
		// appAkademik.Post("/moodle/user/update-password", moodleHandler.UpdatePassword)

		// Fakultas
		fakultasRoute := appAkademik.Group("/fakultas")

		fakultasRoute.Get("/", fakultasHandler.GetFakultas)
		fakultasRoute.Post("/", fakultasHandler.CreateFakultas)
		fakultasRoute.Post("/sync", fakultasHandler.SyncFakultas)

		// Prodi
		prodiRoute := appAkademik.Group("/fakultas/:id/prodi")

		prodiRoute.Get("/", prodiHandler.GetProdis)
		prodiRoute.Post("/", prodiHandler.CreateProdis)
		prodiRoute.Post("/sync", prodiHandler.SyncProdis)

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
	}

	// Postgres Config CRUD
	appSecure.Get("/config/postgres", postgresConfigHandler.ListPostgresConfigs)
	appSecure.Post("/config/postgres", postgresConfigHandler.CreatePostgresConfig)
	appSecure.Get("/config/postgres/:id", postgresConfigHandler.GetPostgresConfig)
	appSecure.Put("/config/postgres/:id", postgresConfigHandler.UpdatePostgresConfig)
	appSecure.Delete("/config/postgres/:id", postgresConfigHandler.DeletePostgresConfig)
}
