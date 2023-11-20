package model

import (
	"context"
)

type UserRepository interface {
	Register(ctx context.Context, user *UserAccounts) (*UserAccounts, error)

	FetchUserByEmail(ctx context.Context, email string) (*UserAccounts, error)
	FetchUserByPhone(ctx context.Context, phone string) (UserAccounts, error)
}

type UserService interface {
	Register(ctx context.Context, user *UserRegisterReq) (*UserRegisterResponse, error)

	FetchUserByEmail(ctx context.Context, email string) (*UserAccounts, error)
	FetchUserByPhone(ctx context.Context, phone string) (LoginUserResponse, error)
	Login(ctx context.Context, user *UserAccounts) (*LoginUserResponse, error)
}
