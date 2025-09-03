package main

import (
	"api/spada/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Register routes
	handler.RegisterRoutes(app)

	app.Listen(":8090")
}
