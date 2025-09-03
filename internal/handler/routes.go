package handler

import (
	"api/spada/internal/service"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {

	userHandler := NewUserHandler(service.NewUserService())
	// kategoriHandler := NewKategoriHandler(service.NewKategoriService(database.DB))

	app.Get("/api/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "pong"})
	})

	// User CRUD
	app.Post("/api/user", userHandler.Index)
	app.Post("/api/user/sync", userHandler.Sync)
	app.Post("/api/user/reset", userHandler.Reset)

	// Kategori CRUD
	// app.Post("/api/kategori", kategoriHandler.Index)
	// app.Post("/api/kategori/sync", kategoriHandler.Sync)
	// app.Post("/api/kategori/reset", kategoriHandler.Reset)
}
