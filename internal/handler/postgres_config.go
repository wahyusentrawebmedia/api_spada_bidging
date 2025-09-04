package handler

import (
	"api/spada/internal/response"
	"api/spada/internal/service"
	"api/spada/internal/utils"
	"strconv"

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
	cc := utils.NewCustomContext(c)

	var req response.PostgresConfigRequest
	if err := c.BodyParser(&req); err != nil {
		return cc.ErrorResponse(err.Error())
	}

	model := req.ToModel()
	config, err := h.Service.Create(nil, &model)
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	return cc.SuccessResponse(config, "Postgres config created successfully")
}

func (h *PostgresConfigHandler) GetPostgresConfig(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return cc.ErrorResponse("Invalid ID")
	}
	config, err := h.Service.GetByID(nil, id)
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}
	return cc.SuccessResponse(config, "Postgres config retrieved successfully")
}

// UpdatePostgresConfig handles updating a Postgres config by ID
func (h *PostgresConfigHandler) UpdatePostgresConfig(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	var req response.PostgresConfigRequest
	if err := c.BodyParser(&req); err != nil {
		return cc.ErrorResponse(err.Error())
	}
	model := req.ToModel()
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return cc.ErrorResponse("Invalid ID")
	}
	model.ID = int(id)
	config, err := h.Service.Update(nil, &model)

	if err != nil {
		return cc.ErrorResponse(err.Error())
	}
	return cc.SuccessResponse(config, "Postgres config updated successfully")
}

// DeletePostgresConfig handles deleting a Postgres config by ID
func (h *PostgresConfigHandler) DeletePostgresConfig(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return cc.ErrorResponse("Invalid ID")
	}
	if err := h.Service.Delete(nil, id); err != nil {
		return cc.ErrorResponse(err.Error())
	}
	return cc.SuccessResponse(nil, "Postgres config deleted successfully")
}

// ListPostgresConfigs handles listing all Postgres configs
func (h *PostgresConfigHandler) ListPostgresConfigs(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	configs, err := h.Service.List(nil)
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	return cc.SuccessResponse(configs, "Postgres configs retrieved successfully")
}
