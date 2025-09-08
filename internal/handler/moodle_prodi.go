package handler

import (
	"api/spada/internal/response"
	"api/spada/internal/service"
	"api/spada/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type MoodleProdisHandler struct {
	service service.MoodleProdiService
}

func NewProdisHandler(service service.MoodleProdiService) *MoodleProdisHandler {
	return &MoodleProdisHandler{service: service}
}

// GET /Prodis
func (h *MoodleProdisHandler) GetProdis(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	id := c.Query("id")

	db, err := cc.GetGormConnectionForPerguruanTinggi()
	if err != nil {
		return cc.ErrorResponse("get database connection: " + err.Error())
	}

	Prodis, err := h.service.GetProdi(id, db)
	if err != nil {
		return cc.ErrorResponse("get Prodis: " + err.Error())
	}

	return cc.SuccessResponse(Prodis, "Prodis fetch successfully")

}

// POST /Prodis
func (h *MoodleProdisHandler) CreateProdis(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	var req response.MoodleProdiRequest
	if err := c.BodyParser(&req); err != nil {
		return cc.ErrorResponse(err.Error())
	}

	db, err := cc.GetGormConnectionForPerguruanTinggi()
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	Prodis, err := h.service.AddProdi(req, db)
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	return cc.SuccessResponse(Prodis, "Prodis created successfully")
}

// POST /Prodis-batch-sync
func (h *MoodleProdisHandler) SyncProdis(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	var req []response.MoodleProdiRequest
	if err := c.BodyParser(&req); err != nil {
		return cc.ErrorResponse(err.Error())
	}

	db, err := cc.GetGormConnectionForPerguruanTinggi()
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	errs := h.service.BatchProdiSync(req, db)
	if errs != nil {
		var errMsgs []string
		for _, err := range errs {
			errMsgs = append(errMsgs, err.Error())
		}
		return cc.ErrorResponse("Batch sync errors: " + utils.JoinStrings(errMsgs, "; "))
	}

	return cc.SuccessResponse(nil, "Prodis sync successfully")
}
