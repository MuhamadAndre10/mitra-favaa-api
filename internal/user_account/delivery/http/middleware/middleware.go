package middleware

import (
	env "github.com/andrepriyanto10/favaa_mitra/config/environment"
	"github.com/andrepriyanto10/favaa_mitra/internal/tool/token"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {

		// get token from header
		authHeader := c.Get("Authorization")

		fields := strings.Fields(authHeader)

		if len(fields) != 2 || fields[0] != "Bearer" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "bad request",
			})
		}

		// load env
		cfg, _ := env.LoadEnv()

		tkn, err := token.ValidateToken(fields[1], cfg.GetString("JWT_SECRET"))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		subject, err := tkn.GetSubject()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "internal server error",
			})
		}

		// save in context
		c.Locals("userId", subject)

		return c.Next()
	}
}
