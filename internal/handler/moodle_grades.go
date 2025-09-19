package handler

import (
	"api/spada/internal/model"
	"api/spada/internal/service"
	"api/spada/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func NewMoodleGradesHandler(service service.MoodleGradesService) *MoodleGradesHandler {
	return &MoodleGradesHandler{service: service}
}

type MoodleGradesHandler struct {
	service service.MoodleGradesService
}

// POST /create-aspek-nilai
func (h *MoodleGradesHandler) CreateAspekNilai(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	var req model.MdlGradeItems
	if err := c.BodyParser(&req); err != nil {
		return cc.ErrorResponse(err.Error())
	}

	db, err := cc.GetGormConnectionForPerguruanTinggi()
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	data, err := h.service.CreateAspekNilai(db, &req)

	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	return cc.SuccessResponse(data, "Aspek Nilai created successfully")
}

// POST /create-aspek-nilai-sync
func (h *MoodleGradesHandler) CreateAspekNilaiSync(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	var req []model.MdlGradeItems
	if err := c.BodyParser(&req); err != nil {
		return cc.ErrorResponse(err.Error())
	}

	db, err := cc.GetGormConnectionForPerguruanTinggi()
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	data, err := h.service.CreateAspekNilaiSync(db, &req)
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}
	return cc.SuccessResponse(data, "Aspek Nilai created successfully")
}
