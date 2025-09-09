package handler

import (
	"api/spada/internal/response"
	"api/spada/internal/service"
	"api/spada/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type MoodledTahunAkademikHandler struct {
	service service.MoodleTahunAkademikService
}

func NewTahunAkademikHandler(service service.MoodleTahunAkademikService) *MoodledTahunAkademikHandler {
	return &MoodledTahunAkademikHandler{service: service}
}

// GET /dTahunAkademik
func (h *MoodledTahunAkademikHandler) GetdTahunAkademik(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	db, err := cc.GetGormConnectionForPerguruanTinggi()
	if err != nil {
		return cc.ErrorResponse("get database connection: " + err.Error())
	}

	dTahunAkademik, err := h.service.GetTahunAkademik(db)
	if err != nil {
		return cc.ErrorResponse("get dTahunAkademik: " + err.Error())
	}

	return cc.SuccessResponse(dTahunAkademik, "dTahunAkademik fetch successfully")

}

// POST /dTahunAkademik
func (h *MoodledTahunAkademikHandler) CreatedTahunAkademik(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	var req response.MoodleTahunAkademikRequest
	if err := c.BodyParser(&req); err != nil {
		return cc.ErrorResponse(err.Error())
	}

	db, err := cc.GetGormConnectionForPerguruanTinggi()
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	dTahunAkademik, err := h.service.AddTahunAkademik(req, db)
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	return cc.SuccessResponse(dTahunAkademik, "dTahunAkademik created successfully")
}

// POST /dTahunAkademik-batch-sync
func (h *MoodledTahunAkademikHandler) SyncdTahunAkademik(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	var req []response.MoodleTahunAkademikRequest
	if err := c.BodyParser(&req); err != nil {
		return cc.ErrorResponse(err.Error())
	}

	db, err := cc.GetGormConnectionForPerguruanTinggi()
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	errs := h.service.BatchTahunAkademikSync(req, db)
	if errs != nil {
		var errMsgs []string
		for _, err := range errs {
			errMsgs = append(errMsgs, err.Error())
		}
		return cc.ErrorResponse("Batch sync errors: " + utils.JoinStrings(errMsgs, "; "))
	}

	return cc.SuccessResponse(nil, "dTahunAkademik sync successfully")
}
