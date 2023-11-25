package repository

import (
	"context"
	"github.com/andrepriyanto10/favaa_mitra/internal/models"

	"gorm.io/gorm"
)

type Repository struct {
	*gorm.DB
}

func NewAuthRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) FetchUserByEmail(ctx context.Context, email string) (*models.UserAccounts, error) {

	var user models.UserAccounts
	err := r.DB.WithContext(ctx).Take(&user, "email = ?", email).Error
	if err != nil {
		return &models.UserAccounts{}, err
	}

	return &user, nil
}

func (r *Repository) FetchUserByPhone(ctx context.Context, phone string) (models.UserAccounts, error) {
	err := r.DB.WithContext(ctx).Take(&models.UserAccounts{}, "phone = ?", phone).Error
	if err != nil {
		return models.UserAccounts{}, err
	}

	return models.UserAccounts{}, nil
}

func (r *Repository) StoreVerificationData(ctx context.Context, data *models.VerificationData) error {
	err := r.DB.WithContext(ctx).Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}
