package handler

import (
	"api/spada/internal/response"
	"api/spada/internal/service"
	"api/spada/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type MoodleMakulHandler struct {
	service *service.MoodleMakulService
}

func NewMoodleMakulHandler(service *service.MoodleMakulService) *MoodleMakulHandler {
	return &MoodleMakulHandler{
		service: service,
	}
}

// POST /moodle/makul/sync
func (h *MoodleMakulHandler) SyncMakul(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	parent := c.Params("semester_id")

	var req []response.MoodleMakulRequest
	if err := c.BodyParser(&req); err != nil {
		return cc.ErrorResponse(err.Error())
	}

	db, err := cc.GetGormConnectionForPerguruanTinggi()
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	errs := h.service.SyncMakulAll(req, parent, db)
	if errs != nil {
		var errMsgs []string
		for _, err := range errs {
			errMsgs = append(errMsgs, err.Error())
		}
		return cc.ErrorResponse("Batch sync errors: " + utils.JoinStrings(errMsgs, "; "))
	}

	return cc.SuccessResponse(nil, "Fakultas sync successfully")
}
