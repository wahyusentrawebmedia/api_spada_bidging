package handler

import (
	"api/spada/internal/model"
	"api/spada/internal/repository"
	"api/spada/internal/service"
	"api/spada/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{UserService: svc}
}

// GET /user
func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	db, err := cc.GetGormConnectionForPerguruanTinggi()
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}
	users, err := h.UserService.FetchAllUsersWithPagination(repository.NewUserRepository(db), 1, 100)
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}
	return cc.SuccessResponse(users, "Users fetched successfully")
}

// POST /user
func (h *UserHandler) Index(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	var req model.UserSyncRequest
	if err := c.BodyParser(&req); err != nil {
		return cc.ErrorResponse(err.Error())
	}

	db, err := cc.GetGormConnectionForPerguruanTinggi()
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	resp, err := h.UserService.SyncUser(cc, repository.NewUserRepository(db), &req)
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}
	return cc.SuccessResponse(resp, "User synced successfully")
}

// // POST /user/sync
// func (h *UserHandler) Sync(c *fiber.Ctx) error {
// 	var req []model.UserSyncRequest
// 	if err := c.BodyParser(&req); err != nil {
// 		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
// 	}
// 	resp := make([]model.UserSyncResponse, 0)
// 	for _, r := range req {
// 		result, _ := h.UserService.SyncUser(&r)
// 		resp = append(resp, *result)
// 	}
// 	return c.JSON(resp)
// }

// // POST /user/reset
// func (h *UserHandler) Reset(c *fiber.Ctx) error {
// 	var ids []int
// 	if err := c.BodyParser(&ids); err != nil {
// 		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
// 	}
// 	count, err := h.UserService.ResetPassword(ids)
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
// 	}
// 	return c.JSON(fiber.Map{"action": true, "jumlah": count})
// }
