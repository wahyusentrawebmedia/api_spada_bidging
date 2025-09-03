package main

import (
	"api/spada/internal/handler"
	"api/spada/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	utils.InitConfig()

	app := fiber.New()

	// Register routes
	handler.RegisterRoutes(app)

	app.Listen(":8090")
}
