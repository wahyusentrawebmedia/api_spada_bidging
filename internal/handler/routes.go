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
		appAkademik.Get("/fakultas", fakultasHandler.GetFakultas)
		appAkademik.Post("/fakultas", fakultasHandler.CreateFakultas)
		appAkademik.Post("/fakultas/sync", fakultasHandler.SyncFakultas)

		// Prodi
		appAkademik.Get("/fakultas/:id/prodi", prodiHandler.GetProdis)
		appAkademik.Post("/fakultas/:id/prodi", prodiHandler.CreateProdis)
		appAkademik.Post("/fakultas/:id/prodi/sync", prodiHandler.SyncProdis)
	}

	// Postgres Config CRUD
	appSecure.Get("/config/postgres", postgresConfigHandler.ListPostgresConfigs)
	appSecure.Post("/config/postgres", postgresConfigHandler.CreatePostgresConfig)
	appSecure.Get("/config/postgres/:id", postgresConfigHandler.GetPostgresConfig)
	appSecure.Put("/config/postgres/:id", postgresConfigHandler.UpdatePostgresConfig)
	appSecure.Delete("/config/postgres/:id", postgresConfigHandler.DeletePostgresConfig)
}
