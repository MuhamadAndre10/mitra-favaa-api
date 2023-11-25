package user_account

import (
	"context"
	"github.com/andrepriyanto10/favaa_mitra/internal/models"
)

type UserAccountService interface {
	Login(ctx context.Context, user *models.UserAccounts) (*LoginUserResponse, error)

	VerificationCode(ctx context.Context, user *UserRegisterReq) error

	FetchUserByEmail(ctx context.Context, email string) (*models.UserAccounts, error)

	StoreVerificationData(ctx context.Context, data *models.VerificationData) error
}
