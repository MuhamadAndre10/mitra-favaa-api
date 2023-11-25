package postgres

import (
	"fmt"
	"github.com/andrepriyanto10/favaa_mitra/internal/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func OpenConnection(cfg *viper.Viper) (*gorm.DB, error) {
	dsnFormat := "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta"

	sslMode := "disable"

	if cfg.GetString("APP_ENV") == "production" {
		sslMode = "require"
	}

	dsn := fmt.Sprintf(dsnFormat,
		cfg.GetString("DB_HOST"),
		cfg.GetString("DB_USERNAME"),
		cfg.GetString("DB_PASSWORD"),
		cfg.GetString("DB_DATABASE"),
		cfg.GetString("DB_PORT"),
		sslMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		return nil, err
	}

	log.Println("Database connection established successfully!")

	log.Println("Migrating database...")

	err = db.Migrator().DropTable(&models.UserAccounts{}, &models.Partner{}, &models.Address{}, &models.VerificationData{})
	if err != nil {
		return nil, err
	}
	log.Println("Database dropped successfully!")

	err = db.AutoMigrate(&models.UserAccounts{}, &models.Partner{}, &models.Address{}, &models.VerificationData{})

	if err != nil {
		return nil, err
	}

	log.Println("Database migrated successfully!")

	return db, nil
}
