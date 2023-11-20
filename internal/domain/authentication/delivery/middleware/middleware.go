package middleware

import (
	"github.com/andrepriyanto10/favaa_mitra/internal/configs/environment"
	token_jwt "github.com/andrepriyanto10/favaa_mitra/internal/configs/jwt"
	"github.com/gofiber/fiber/v2"
	"strings"
)

// Protected make middleware for protected route

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
		env, _ := environment.LoadEnv()

		token, err := token_jwt.ValidateToken(fields[1], env.GetString("JWT_SECRET"))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		subject, err := token.GetSubject()
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
