package main

import (
	"fmt"
	"github.com/andrepriyanto10/favaa_mitra/config/database/postgres"
	"github.com/andrepriyanto10/favaa_mitra/config/environment"
	"github.com/andrepriyanto10/favaa_mitra/internal/user_account/delivery/http"
	"github.com/andrepriyanto10/favaa_mitra/internal/user_account/repository"
	"github.com/andrepriyanto10/favaa_mitra/internal/user_account/service"
	"github.com/andrepriyanto10/favaa_mitra/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const idleTimeout = 5 * time.Second

func main() {

	app := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
	})

	app.Use(logger.New(logger.Config{
		Format: "${pid} ${status} - ${method} ${path}\n",
	}))

	app.Use(cors.New())

	cfg, err := env.env.LoadEnv()
	if err != nil {
		log.Fatalf("Error Loading environment: %v", err)
	}

	db, err := postgres.OpenConnection(cfg)
	if err != nil {
		log.Panic("fail connect", err)
	}

	authRepo := repository.NewAuthRepository(db)

	authService := service.NewAuthService(authRepo, cfg)

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
