package middlewares

import (
	"nls-auth/api/v1/auth/db"
	"nls-auth/constants"

	"github.com/gofiber/fiber/v2"
)

func RequestDBMiddleware(c *fiber.Ctx) error {
	//application := c.Context().Value(constants.ApplicationCtx).(constants.AppKey)
	DB := c.Locals(constants.DBLocals).(*db.PG_DB)

	if c != nil {
		c.Locals(constants.RequestDBLocals, DB)
	} else {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}
	return c.Next()
}
