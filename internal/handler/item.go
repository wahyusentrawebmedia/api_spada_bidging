package handler

import (
	"api/spada/internal/model"

	"github.com/gofiber/fiber/v2"
)

var items = []model.Item{}

func GetItems(c *fiber.Ctx) error {
	return c.JSON(items)
}

func CreateItem(c *fiber.Ctx) error {
	item := new(model.Item)
	if err := c.BodyParser(item); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	items = append(items, *item)
	return c.Status(201).JSON(item)
}

func GetItem(c *fiber.Ctx) error {
	id := c.Params("id")
	for _, item := range items {
		if item.ID == id {
			return c.JSON(item)
		}
	}
	return c.Status(404).JSON(fiber.Map{"error": "Item not found"})
}

func UpdateItem(c *fiber.Ctx) error {
	id := c.Params("id")
	for i, item := range items {
		if item.ID == id {
			updated := new(model.Item)
			if err := c.BodyParser(updated); err != nil {
				return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
			}
			items[i] = *updated
			return c.JSON(updated)
		}
	}
	return c.Status(404).JSON(fiber.Map{"error": "Item not found"})
}

func DeleteItem(c *fiber.Ctx) error {
	id := c.Params("id")
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			return c.SendStatus(204)
		}
	}
	return c.Status(404).JSON(fiber.Map{"error": "Item not found"})
}
