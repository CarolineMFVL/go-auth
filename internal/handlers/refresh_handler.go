package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func RefreshHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Refresh token logic not implemented yet",
	})
}
