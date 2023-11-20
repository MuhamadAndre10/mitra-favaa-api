package main

import (
	"fmt"
	handler "github.com/andrepriyanto10/favaa_mitra/internal/domain/authentication/delivery/http"
	"github.com/andrepriyanto10/favaa_mitra/internal/domain/authentication/repository"
	"github.com/andrepriyanto10/favaa_mitra/internal/domain/authentication/services"
	"github.com/andrepriyanto10/favaa_mitra/internal/router"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	db, app := setup()

	authRepo := repository.NewAuthRepository(db)

	authService := services.NewAuthService(authRepo)

	authHandler := handler.NewAuthHandler(authService)

	// init routes
	router.InitRoutes(app, authHandler)

	// run server
	shutdown(app)
}

func shutdown(app *fiber.App) {
	go func() {
		err := app.Listen(":8000")
		if err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	_ = <-c

	fmt.Println("Gracefully shutting down...")

	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// db.Close()

	fmt.Println("Done.")
}
