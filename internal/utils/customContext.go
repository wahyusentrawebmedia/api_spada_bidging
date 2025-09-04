package utils

import (
	"api/spada/internal/response"

	"github.com/gofiber/fiber/v2"
)

type CustomContext struct {
	*fiber.Ctx
	// Add custom fields here if needed, e.g. UserID, RequestID, etc.
}

func NewCustomContext(c *fiber.Ctx) *CustomContext {
	return &CustomContext{Ctx: c}
}

func (cc *CustomContext) SuccessResponse(data interface{}, message string) error {
	return cc.Status(fiber.StatusOK).JSON(response.DefaultResponse{
		Status:  true,
		Message: message,
		Data:    data,
	})
}

func (cc *CustomContext) ErrorResponse(message string, code int) error {
	return cc.Status(code).JSON(response.DefaultResponse{
		Status:  false,
		Message: message,
	})
}
