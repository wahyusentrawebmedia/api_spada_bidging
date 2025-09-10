package handler

import (
	"api/spada/internal/response"
	"api/spada/internal/service"
	"api/spada/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type MoodleCategoriesHandler struct {
	service service.MoodleCategoriesService
}

func NewCategoriesHandler(service service.MoodleCategoriesService) *MoodleCategoriesHandler {
	return &MoodleCategoriesHandler{service: service}
}

// GET /Categories
func (h *MoodleCategoriesHandler) GetCategories(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	db, err := cc.GetGormConnectionForPerguruanTinggi()
	if err != nil {
		return cc.ErrorResponse("get database connection: " + err.Error())
	}

	Categories, err := h.service.GetCategories(db)
	if err != nil {
		return cc.ErrorResponse("get Categories: " + err.Error())
	}

	return cc.SuccessResponse(Categories, "Categories fetch successfully")

}

// POST /Categories
func (h *MoodleCategoriesHandler) CreateCategories(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	var req response.MoodleCategoriesRequest
	if err := c.BodyParser(&req); err != nil {
		return cc.ErrorResponse(err.Error())
	}

	db, err := cc.GetGormConnectionForPerguruanTinggi()
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	Categories, err := h.service.AddCategories(req, db)
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	return cc.SuccessResponse(Categories, "Categories created successfully")
}

// POST /Categories-batch-sync
func (h *MoodleCategoriesHandler) SyncCategories(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	var req []response.MoodleCategoriesRequest
	if err := c.BodyParser(&req); err != nil {
		return cc.ErrorResponse(err.Error())
	}

	db, err := cc.GetGormConnectionForPerguruanTinggi()
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	errs := h.service.BatchCategoriesSync(req, db)
	if errs != nil {
		var errMsgs []string
		for _, err := range errs {
			errMsgs = append(errMsgs, err.Error())
		}
		return cc.ErrorResponse("Batch sync errors: " + utils.JoinStrings(errMsgs, "; "))
	}

	return cc.SuccessResponse(nil, "Categories sync successfully")
}

// GET /Categories-with-prefix
func (h *MoodleCategoriesHandler) GetCategoriesWithPrefix(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	prefix := c.Params("prefix")
	backStr := c.Query("back")

	back := false
	if backStr == "true" || backStr == "1" {
		back = true
	}

	prefix = utils.CleanURLParam(prefix)
	prefix = utils.ReplaceAll(prefix, "%20", " ")

	db, err := cc.GetGormConnectionForPerguruanTinggi()
	if err != nil {
		return cc.ErrorResponse("get database connection: " + err.Error())
	}

	Categories, err := h.service.GetCategoriesWithPrefix(prefix, back, db)
	if err != nil {
		return cc.ErrorResponse("get Categories with prefix: " + err.Error())
	}

	return cc.SuccessResponse(Categories, "Categories with prefix fetch successfully")

}
