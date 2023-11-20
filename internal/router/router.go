package router

import (
	handler "github.com/andrepriyanto10/favaa_mitra/internal/domain/authentication/delivery/http"
	"github.com/andrepriyanto10/favaa_mitra/internal/domain/authentication/delivery/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func InitRoutes(app *fiber.App, authHandler handler.AuthHandler) {
	app.Use(cors.New())

	// Changing TimeZone & TimeFormat
	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "America/New_York",
	}))

	//app.Use("/api/v1", func(c *fiber.Ctx) error {
	//	return c.Next()
	//})

	app.Get("/hello", middleware.Protected(), func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Post("/register", authHandler.Register)
	app.Post("/login", authHandler.Login)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Page not found",
		}) // => 404 "Not Found"
	})

}
