package user_account

import (
	"context"
	"github.com/andrepriyanto10/favaa_mitra/internal/models"
)

type UserAccountRepository interface {
	FetchUserByEmail(ctx context.Context, email string) (*models.UserAccounts, error)
	StoreVerificationData(ctx context.Context, data *models.VerificationData) error
}
