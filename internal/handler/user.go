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

	parameter := service.ParameterUser{
		IdNumberGroup: c.Query("id_number_group"),
		TypeUser:      c.Query("type_user"),
		IdMakul:       c.Query("kode_makul"),
		Page:          1,
		Limit:         100,
	}

	users, err := h.UserService.FetchAllUsersWithPagination(
		db,
		parameter,
	)
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}
	return cc.SuccessResponse(users, "Users fetched successfully")
}

// POST /user
func (h *UserHandler) UpdateSingle(c *fiber.Ctx) error {
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

// GetDetail /user
func (h *UserHandler) GetDetail(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	db, err := cc.GetGormConnectionForPerguruanTinggi()
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	user, err := h.UserService.GetUserByUsername(repository.NewUserRepository(db), cc.GetUsername())

	if err != nil {
		return cc.ErrorResponse(err.Error())
	}
	return cc.SuccessResponse(user, "User fetched successfully")
}

// UpdatePassword /user/password
func (h *UserHandler) UpdatePassword(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	var req model.UserChangePasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return cc.ErrorResponse(err.Error())
	}

	db, err := cc.GetGormConnectionForPerguruanTinggi()
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	err = h.UserService.ChangePassword(cc, repository.NewUserRepository(db), cc.GetUsername(), req.OldPassword, req.NewPassword)
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}
	return cc.SuccessResponse(fiber.Map{"action": true}, "Password updated successfully")
}

// POST /user/sync
func (h *UserHandler) Sync(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	var req []model.UserSyncRequest
	if err := c.BodyParser(&req); err != nil {
		return cc.ErrorResponse(err.Error())
	}

	db, err := cc.GetGormConnectionForPerguruanTinggi()
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	resp, err := h.UserService.SyncUserBatch(cc, repository.NewUserRepository(db), req)
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}
	return cc.SuccessResponse(resp, "Users synced successfully")
}

// POST /user/sync/dosen-mahasiswa
func (h *UserHandler) SyncDosenMahasiswa(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	var req model.DosenMahasiwaSyncRequest
	if err := c.BodyParser(&req); err != nil {
		return cc.ErrorResponse(err.Error())
	}

	db, err := cc.GetGormConnectionForPerguruanTinggi()
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	resp, err := h.UserService.SyncUserBatchDosenMahasiswa(cc, db, req)
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}
	return cc.SuccessResponse(resp, "Users synced successfully")
}

// POST /user/sync/dosen-mahasiswa
func (h *UserHandler) SyncDosenMahasiswaMakul(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	kodeMakul := c.Params("id_makul")

	var req model.DosenMahasiwaSyncRequest
	if err := c.BodyParser(&req); err != nil {
		return cc.ErrorResponse(err.Error())
	}

	db, err := cc.GetGormConnectionForPerguruanTinggi()
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	resp, err := h.UserService.SyncUserBatchDosenMahasiswaMakul(cc, db, req, service.DosenMahasiwaSyncRequest{
		KodeMakul: kodeMakul,
	})

	if err != nil {
		return cc.ErrorResponse(err.Error())
	}
	return cc.SuccessResponse(resp, "Users synced successfully")
}

// POST /user/sync/dosen-mahasiswa
func (h *UserHandler) SyncDosenMahasiswaCategories(c *fiber.Ctx) error {
	cc := utils.NewCustomContext(c)

	kodeCategories := c.Params("kode_categories")

	var req model.DosenMahasiwaSyncRequest
	if err := c.BodyParser(&req); err != nil {
		return cc.ErrorResponse(err.Error())
	}

	db, err := cc.GetGormConnectionForPerguruanTinggi()
	if err != nil {
		return cc.ErrorResponse(err.Error())
	}

	resp, err := h.UserService.SyncUserBatchDosenMahasiswaMakul(cc, db, req, service.DosenMahasiwaSyncRequest{
		KodeCategories: kodeCategories,
	})

	if err != nil {
		return cc.ErrorResponse(err.Error())
	}
	return cc.SuccessResponse(resp, "Users synced successfully")
}

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
