package main

import (
	"github.com/andrepriyanto10/favaa_mitra/internal/configs/database"
	"github.com/andrepriyanto10/favaa_mitra/internal/configs/environment"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"time"
)

const idleTimeout = 5 * time.Second

func setup() (db *gorm.DB, app *fiber.App) {
	// load env
	env, err := environment.LoadEnv()
	if err != nil {
		log.Fatalln("Error Loading environment", err)
	}

	// connection db
	db, err = database.OpenConnection(env)
	if err != nil {
		log.Panic("fail connect", err)
	}

	// use fiber framework for app
	app = fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
	})

	return db, app
}
