package repository

import (
	"context"
	"github.com/andrepriyanto10/favaa_mitra/internal/domain/authentication/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	gorm.DB
}

func NewAuthRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: *db}
}

func (r *UserRepository) Register(ctx context.Context, user *model.UserAccounts) (*model.UserAccounts, error) {
	err := r.DB.WithContext(ctx).Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FetchUserByEmail(ctx context.Context, email string) (*model.UserAccounts, error) {

	var user model.UserAccounts
	err := r.DB.WithContext(ctx).Take(&user, "email = ?", email).Error
	if err != nil {
		return &model.UserAccounts{}, err
	}

	return &user, nil
}

func (r *UserRepository) FetchUserByPhone(ctx context.Context, phone string) (model.UserAccounts, error) {
	err := r.DB.WithContext(ctx).Take(&model.UserAccounts{}, "phone = ?", phone).Error
	if err != nil {
		return model.UserAccounts{}, err
	}

	return model.UserAccounts{}, nil
}
