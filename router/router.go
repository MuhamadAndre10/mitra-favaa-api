package router

import (
	handler "github.com/andrepriyanto10/favaa_mitra/internal/user_account/delivery/http"
	"github.com/andrepriyanto10/favaa_mitra/internal/user_account/delivery/http/middleware"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App, authHandler *handler.AuthHandler) {

	app.Get("/hello", middleware.Protected(), func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Post("/verification-code", authHandler.VerificationCode)
	app.Post("/login", authHandler.Login)
	app.Post("/send-otp", authHandler.GeneratePassResetCode)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Page not found",
		}) // => 404 "Not Found"
	})

}
