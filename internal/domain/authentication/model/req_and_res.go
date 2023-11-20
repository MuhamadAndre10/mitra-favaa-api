package model

import "time"

type UserRegisterReq struct {
	Username string `json:"username,omitempty" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email"`
	Phone    string `json:"phone,omitempty"  validate:"required"`
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
