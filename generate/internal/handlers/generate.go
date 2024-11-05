package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GenerateHandler(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}
