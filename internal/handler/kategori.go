package handler

import (
	"api/spada/internal/model"
	"api/spada/internal/service"

	"github.com/gofiber/fiber/v2"
)

type KategoriHandler struct {
	service *service.KategoriService
}

func NewKategoriHandler(svc *service.KategoriService) *KategoriHandler {
	return &KategoriHandler{service: svc}
}

func (h *KategoriHandler) Index(c *fiber.Ctx) error {
	var req model.RequestKategori
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"action": false, "error": "Invalid request"})
	}
	data, err := h.service.GetAllKategori(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"action": false, "error": err.Error()})
	}
	return c.JSON(fiber.Map{"action": true, "data": data})
}

// func (h *KategoriHandler) Sinkron(c *fiber.Ctx) error {
// 	var req struct {
// 		Tahun string `json:"tahun"`
// 		Name  string `json:"name"`
// 	}
// 	if err := c.BodyParser(&req); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"action": false, "error": "Invalid request"})
// 	}
// 	result, err := h.service.SinkronKategori(req.Tahun, req.Name)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"action": true, "aksi": "gagal", "pesan": err.Error()})
// 	}
// 	return c.JSON(result)
// }

// func (h *KategoriHandler) Prodi(c *fiber.Ctx) error {
// 	var records []struct {
// 		IdFak    int    `json:"id_fak"`
// 		Name     string `json:"name"`
// 		IdNumber string `json:"idnumber"`
// 	}
// 	if err := c.BodyParser(&records); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"action": false, "error": "Invalid request"})
// 	}
// 	resp, err := h.service.HandleProdi(records)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"action": false, "error": err.Error()})
// 	}
// 	return c.JSON(resp)
// }

// func (h *KategoriHandler) Fakultas(c *fiber.Ctx) error {
// 	var records []struct {
// 		Name     string `json:"name"`
// 		IdNumber string `json:"idnumber"`
// 	}
// 	if err := c.BodyParser(&records); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"action": false, "error": "Invalid request"})
// 	}
// 	resp, err := h.service.HandleFakultas(records)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"action": false, "error": err.Error()})
// 	}
// 	return c.JSON(resp)
// }
