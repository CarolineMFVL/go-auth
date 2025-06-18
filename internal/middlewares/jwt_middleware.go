package middlewares

import (

	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware(c *fiber.Ctx) error {
	// Vérification du token JWT
	return c.Next()
}
