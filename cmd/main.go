package main

import (
	"api/spada/internal/handler"
	"api/spada/internal/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	utils.InitConfig()

	app := fiber.New()

	// Rate limiting khusus untuk API endpoints
	apiLimiter := limiter.New(limiter.Config{
		Max:        10,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
	})

	app.Use(apiLimiter)

	// Register routes
	handler.RegisterRoutes(app)

	if err := app.Listen(":8100"); err != nil {
		panic(err)
	}
}
