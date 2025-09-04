package handler

import (
	"api/spada/internal/database"
	"api/spada/internal/repository"
	"api/spada/internal/service"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	repoKategori := repository.NewPostgresConfigRepository(database.DB)
	servicesConfig := service.NewPostgresConfigService(*repoKategori)

	postgresConfigHandler := NewPostgresConfigHandler(*servicesConfig)
	userHandler := NewUserHandler(service.NewUserService())

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "pong"})
	})

	// User CRUD
	app.Post("/user", userHandler.Index)
	app.Post("/user/sync", userHandler.Sync)
	app.Post("/user/reset", userHandler.Reset)

	// Postgres Config CRUD
	app.Get("/config/postgres", postgresConfigHandler.ListPostgresConfigs)
	app.Post("/config/postgres", postgresConfigHandler.CreatePostgresConfig)
	app.Get("/config/postgres/:id", postgresConfigHandler.GetPostgresConfig)
	app.Put("/config/postgres/:id", postgresConfigHandler.UpdatePostgresConfig)
	app.Delete("/config/postgres/:id", postgresConfigHandler.DeletePostgresConfig)
}
