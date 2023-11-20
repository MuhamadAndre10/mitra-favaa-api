package authentication

import (
	"context"
	"github.com/andrepriyanto10/favaa_mitra/internal/domain_impl/authentication/utils"
)

type UserRepository interface {
	Register(ctx context.Context, user *UserAccounts) (*UserAccounts, error)
	FetchUserByEmail(ctx context.Context, email string) (*UserAccounts, error)
	FetchUserByPhone(ctx context.Context, phone string) (UserAccounts, error)
}

type UserService interface {
	Register(ctx context.Context, user *utils.UserRegisterReq) (*utils.UserRegisterResponse, error)
	FetchUserByEmail(ctx context.Context, email string) (*UserAccounts, error)
	FetchUserByPhone(ctx context.Context, phone string) (utils.LoginUserResponse, error)
	Login(ctx context.Context, user *UserAccounts) (*utils.LoginUserResponse, error)
}
