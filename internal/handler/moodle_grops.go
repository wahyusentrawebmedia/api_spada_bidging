package handler

import (
	"api/spada/internal/service"
	"api/spada/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type GroupsHandler struct {
	service service.MoodleGroupsService
}

func NewGroupsHandler(service service.MoodleGroupsService) *GroupsHandler {
	return &GroupsHandler{service: service}
}

// GET /groups/categories/:categories_id
func (h *GroupsHandler) GetGroupsByCategoriesID(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	categoriesID := c.Params("categories_id")

	categoriesID = utils.ReplaceAll(categoriesID, "%20", " ")

	db, err := cc.GetGormConnectionForPerguruanTinggi()
	if err != nil {
		return cc.ErrorResponse("get database connection: " + err.Error())
	}

	groups, err := h.service.GetGroupsByCategoriesID(categoriesID, db)
	if err != nil {
		return cc.ErrorResponse("get groups by categories ID: " + err.Error())
	}

	return cc.SuccessResponse(groups, "Groups fetched successfully")
}
