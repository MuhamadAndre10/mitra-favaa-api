package utils

import "github.com/gofiber/fiber/v2"

func Response(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"message": message,
		"data":    data,
	})
}
