package handler

import (
	"api/spada/internal/response"
	"api/spada/internal/service"

	"github.com/gofiber/fiber/v2"
) // CreatePostgresConfig handles creating a new Postgres config

type PostgresConfigHandler struct {
	Service service.PostgresConfigService
}

func NewPostgresConfigHandler(svc service.PostgresConfigService) *PostgresConfigHandler {
	return &PostgresConfigHandler{Service: svc}
}

// CreatePostgresConfig handles creating a new Postgres config
func (h *PostgresConfigHandler) CreatePostgresConfig(c *fiber.Ctx) error {
	var req response.PostgresConfigRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	config, err := h.Service.Create(c, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(config)
}

// GetPostgresConfig handles fetching a Postgres config by ID
func (h *PostgresConfigHandler) GetPostgresConfig(c *fiber.Ctx) error {
	id := c.Params("id")
	config, err := h.Service.GetByID()(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(config)
}

// UpdatePostgresConfig handles updating a Postgres config by ID
func (h *PostgresConfigHandler) UpdatePostgresConfig(c *fiber.Ctx) error {
	id := c.Params("id")
	var req response.PostgresConfigRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	config, err := h.Service.Update(id, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(config)
}

// DeletePostgresConfig handles deleting a Postgres config by ID
func (h *PostgresConfigHandler) DeletePostgresConfig(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.Service.Delete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// ListPostgresConfigs handles listing all Postgres configs
func (h *PostgresConfigHandler) ListPostgresConfigs(c *fiber.Ctx) error {
	configs, err := h.Service.List()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(configs)
}
