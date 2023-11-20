package repository

import (
	"context"
	"github.com/andrepriyanto10/favaa_mitra/internal/domain/authentication"
	"gorm.io/gorm"
)

type UserRepository struct {
	gorm.DB
}

func NewAuthRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: *db}
}

func (r *UserRepository) Register(ctx context.Context, user *authentication.UserAccounts) (*authentication.UserAccounts, error) {
	err := r.DB.WithContext(ctx).Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FetchUserByEmail(ctx context.Context, email string) (*authentication.UserAccounts, error) {

	var user authentication.UserAccounts
	err := r.DB.WithContext(ctx).Take(&user, "email = ?", email).Error
	if err != nil {
		return &authentication.UserAccounts{}, err
	}

	return &user, nil
}

func (r *UserRepository) FetchUserByPhone(ctx context.Context, phone string) (authentication.UserAccounts, error) {
	err := r.DB.WithContext(ctx).Take(&authentication.UserAccounts{}, "phone = ?", phone).Error
	if err != nil {
		return authentication.UserAccounts{}, err
	}

	return authentication.UserAccounts{}, nil
}
