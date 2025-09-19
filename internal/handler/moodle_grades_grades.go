package handler

import (
	"api/spada/internal/model"
	"api/spada/internal/service"
	"api/spada/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func NewMoodleGradesGradesItemsHandler(service service.MoodleGradesItemsService) *MoodleGradesItemsHandler {
	return &MoodleGradesItemsHandler{service: service}
}

type MoodleGradesItemsHandler struct {
	service service.MoodleGradesItemsService
}

// POST /create-grade-items
func (h *MoodleGradesItemsHandler) CreateGradeItems(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	var req model.GradeGrade
	if err := c.BodyParser(&req); err != nil {
		return cc.ErrorResponse(err.Error())
	}

	db, err := cc.GetGormConnectionForPerguruanTinggi()
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	data, err := h.service.CreateGradeItems(db, &req)

	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	return cc.SuccessResponse(data, "Grade Items created successfully")
}

// POST /create-grade-items-sync
func (h *MoodleGradesItemsHandler) CreateGradeItemsSync(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	var req []model.GradeGrade
	if err := c.BodyParser(&req); err != nil {
		return cc.ErrorResponse(err.Error())
	}

	db, err := cc.GetGormConnectionForPerguruanTinggi()
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	data, err := h.service.CreateGradeItemsSync(db, &req)
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}
	return cc.SuccessResponse(data, "Grade Items created successfully")
}
