package services

import (
	"context"
	"github.com/andrepriyanto10/favaa_mitra/internal/configs/environment"
	token_jwt "github.com/andrepriyanto10/favaa_mitra/internal/configs/jwt"
	"github.com/andrepriyanto10/favaa_mitra/internal/domain/authentication/model"
	"github.com/andrepriyanto10/favaa_mitra/utils"
	"net/http"
	"strings"
	"time"
)

type AuthService struct {
	model.UserRepository
	timeout time.Duration
}

func NewAuthService(user model.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: user,
		timeout:        time.Duration(2) * time.Second,
	}
}

func (u *AuthService) Register(ctx context.Context, user *model.UserRegisterReq) (*model.UserRegisterResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	if user.Username == "" || user.Email == "" || user.Password == "" {
		return nil, &utils.CustomError{Code: http.StatusBadRequest, Message: "Bad Request"}
	}

	if len(user.Username) < 5 || len(user.Username) > 32 {
		return nil, &utils.CustomError{Code: http.StatusBadRequest, Message: "Username must be between 5 and 32 characters"}
	}

	// validation email and sending error response
	if !strings.Contains(strings.TrimSpace(user.Email), "@") {
		return nil, &utils.CustomError{Code: http.StatusBadRequest, Message: "Invalid email"}
	}
	// validation password and sending error response
	if len(user.Password) < 8 || len(user.Password) > 32 {
		return nil, &utils.CustomError{Code: http.StatusBadRequest, Message: "Password must be between 8 and 32 characters"}
	}

	_, err := u.UserRepository.FetchUserByEmail(ctx, user.Email)
	if err == nil {
		return nil, &utils.CustomError{Code: http.StatusBadRequest, Message: "Email already registered"}
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, &utils.CustomError{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}

	newUser := &model.UserAccounts{
		Username:       user.Username,
		Email:          user.Email,
		Phone:          user.Phone,
		Password:       hashedPassword,
		Status:         "active",
		DateRegistered: time.Now().UTC(),
	}

	userRegistered, err := u.UserRepository.Register(ctx, newUser)
	if err != nil {
		return nil, &utils.CustomError{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}

	registerRes := &model.UserRegisterResponse{
		ID:             userRegistered.ID,
		Email:          userRegistered.Email,
		Phone:          userRegistered.Phone,
		DateRegistered: userRegistered.DateRegistered,
	}

	return registerRes, nil

}

func (u *AuthService) FetchUserByEmail(ctx context.Context, email string) (*model.UserAccounts, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	if !utils.Valid(email) {
		return nil, &utils.CustomError{Code: http.StatusBadRequest, Message: "Bad Request"}
	}

	if email == "" || !strings.Contains(strings.TrimSpace(email), "@") {
		return nil, &utils.CustomError{Code: http.StatusBadRequest, Message: "Bad Request"}
	}

	user, err := u.UserRepository.FetchUserByEmail(ctx, email)
	if err != nil {
		return nil, &utils.CustomError{Code: http.StatusInternalServerError, Message: "User not found"}
	}

	return user, nil
}

func (u *AuthService) FetchUserByPhone(ctx context.Context, phone string) (model.LoginUserResponse, error) {
	panic("implement me")
}

func (u *AuthService) Login(ctx context.Context, user *model.UserAccounts) (*model.LoginUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	env, _ := environment.LoadEnv()

	tokenExpired, _ := time.ParseDuration(env.GetString("JWT_TTL"))

	// token jwt
	token, err := token_jwt.GenerateToken(tokenExpired, user.ID, env.GetString("JWT_SECRET"))
	if err != nil {
		return nil, &utils.CustomError{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}

	loginRes := &model.LoginUserResponse{
		Token: token,
	}

	return loginRes, nil

}
