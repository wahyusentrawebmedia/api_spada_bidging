package handler

import (
	"api/spada/internal/response"
	"api/spada/internal/service"
	"api/spada/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type MoodleFakultasHandler struct {
	service service.MoodleFakultasService
}

func NewFakultasHandler(service service.MoodleFakultasService) *MoodleFakultasHandler {
	return &MoodleFakultasHandler{service: service}
}

// GET /fakultas
func (h *MoodleFakultasHandler) GetFakultas(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)
	var req response.MoodleFakultasRequest
	if err := c.BodyParser(&req); err != nil {
		return cc.ErrorResponse(err.Error())
	}

	db, err := cc.GetGormConnectionForPerguruanTinggi()
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	fakultas, err := h.service.GetFakultas(req, db)
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	return cc.SuccessResponse(fakultas, "Fakultas fetch successfully")

}

// POST /fakultas
func (h *MoodleFakultasHandler) CreateFakultas(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	var req response.MoodleFakultasRequest
	if err := c.BodyParser(&req); err != nil {
		return cc.ErrorResponse(err.Error())
	}

	db, err := cc.GetGormConnectionForPerguruanTinggi()
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	fakultas, err := h.service.AddFakultas(req, db)
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	return cc.SuccessResponse(fakultas, "Fakultas created successfully")
}
