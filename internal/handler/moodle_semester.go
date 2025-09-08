package handler

import (
	"api/spada/internal/response"
	"api/spada/internal/service"
	"api/spada/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type MoodledSemesterHandler struct {
	service service.MoodleSemesterService
}

func NewSemesterHandler(service service.MoodleSemesterService) *MoodledSemesterHandler {
	return &MoodledSemesterHandler{service: service}
}

// GET /dSemester
func (h *MoodledSemesterHandler) GetdSemester(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	db, err := cc.GetGormConnectionForPerguruanTinggi()
	if err != nil {
		return cc.ErrorResponse("get database connection: " + err.Error())
	}

	dSemester, err := h.service.GetSemester(db)
	if err != nil {
		return cc.ErrorResponse("get dSemester: " + err.Error())
	}

	return cc.SuccessResponse(dSemester, "dSemester fetch successfully")

}

// POST /dSemester
func (h *MoodledSemesterHandler) CreatedSemester(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	var req response.MoodleSemesterRequest
	if err := c.BodyParser(&req); err != nil {
		return cc.ErrorResponse(err.Error())
	}

	db, err := cc.GetGormConnectionForPerguruanTinggi()
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	dSemester, err := h.service.AddSemester(req, db)
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	return cc.SuccessResponse(dSemester, "dSemester created successfully")
}

// POST /dSemester-batch-sync
func (h *MoodledSemesterHandler) SyncdSemester(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	var req []response.MoodleSemesterRequest
	if err := c.BodyParser(&req); err != nil {
		return cc.ErrorResponse(err.Error())
	}

	db, err := cc.GetGormConnectionForPerguruanTinggi()
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	errs := h.service.BatchSemesterSync(req, db)
	if errs != nil {
		var errMsgs []string
		for _, err := range errs {
			errMsgs = append(errMsgs, err.Error())
		}
		return cc.ErrorResponse("Batch sync errors: " + utils.JoinStrings(errMsgs, "; "))
	}

	return cc.SuccessResponse(nil, "dSemester sync successfully")
}
