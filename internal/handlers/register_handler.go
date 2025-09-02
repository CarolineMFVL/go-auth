package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func RegisterHandler(c *fiber.Ctx) error {
	json.NewEncoder(c).Encode(map[string]string{"message": "Utilisateur créé"})
	return nil
}
