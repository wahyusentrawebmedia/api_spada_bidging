package handler

import (
	"api/spada/internal/model"
	"api/spada/internal/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{UserService: svc}
}

// POST /user
func (h *UserHandler) Index(c *fiber.Ctx) error {
	var req model.UserSyncRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	resp, err := h.UserService.SyncUser(&req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}

// POST /user/sync
func (h *UserHandler) Sync(c *fiber.Ctx) error {
	var req []model.UserSyncRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	resp := make([]model.UserSyncResponse, 0)
	for _, r := range req {
		result, _ := h.UserService.SyncUser(&r)
		resp = append(resp, *result)
	}
	return c.JSON(resp)
}

// POST /user/reset
func (h *UserHandler) Reset(c *fiber.Ctx) error {
	var ids []int
	if err := c.BodyParser(&ids); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	count, err := h.UserService.ResetPassword(ids)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"action": true, "jumlah": count})
}
