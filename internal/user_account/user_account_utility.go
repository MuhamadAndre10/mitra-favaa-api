package user_account

import "time"

type UserRegisterReq struct {
	Username string `json:"username,omitempty" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required,min=8,max=32"`
}

type UserRegisterResponse struct {
	ID             string    `json:"id,omitempty" `
	Email          string    `json:"email,omitempty"`
	Phone          string    `json:"phone,omitempty"`
	DateRegistered time.Time `json:"date_registered"`
}

type LoginUserRequest struct {
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required,min=8,max=32"`
}

type LoginUserResponse struct {
	Token string `json:"token,omitempty"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email,omitempty" validate:"required,email"`
}

type ForgotPasswordResponse struct {
	Email   string `json:"email,omitempty"`
	Message string `json:"message,omitempty"`
}
