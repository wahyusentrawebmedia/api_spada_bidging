package handler

import (
	"api/spada/internal/service"

	"github.com/gofiber/fiber/v2"
)

type MoodleHandler struct {
	// Service *service.MoodleService
}

func NewMoodleHandler(svc *service.MoodleService) *MoodleHandler {
	return &MoodleHandler{}
}

// POST /moodle/user/update-password
func (h *MoodleHandler) UpdatePassword(c *fiber.Ctx) error {
	type reqBody struct {
		UserID   int    `json:"user_id"`
		Password string `json:"password"`
	}
	var req reqBody
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	// if err := h.Service.UpdateUserPassword(req.UserID, req.Password); err != nil {
	// 	return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	// }
	return c.JSON(fiber.Map{"success": true})
}
