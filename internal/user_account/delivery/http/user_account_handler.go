package handler

import (
	"errors"
	"github.com/andrepriyanto10/favaa_mitra/internal/helpers"
	"github.com/andrepriyanto10/favaa_mitra/internal/user_account"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type AuthHandler struct {
	user_account.UserAccountService
}

func NewAuthHandler(service user_account.UserAccountService) *AuthHandler {
	return &AuthHandler{
		UserAccountService: service,
	}

}

func (h *AuthHandler) VerificationCode(c *fiber.Ctx) error {
	if c.Method() != http.MethodPost {
		return helpers.Response(c, http.StatusMethodNotAllowed, "Method Not Allowed", nil)
	}

	userRequest := new(user_account.UserRegisterReq)
	var customErr *helpers.CustomError

	// body parsing
	if err := c.BodyParser(userRequest); err != nil {
		return helpers.Response(c, http.StatusBadRequest, "Bad Request", nil)
	}

	err := h.UserAccountService.VerificationCode(c.Context(), userRequest)

	if errors.As(err, &customErr) {
		return helpers.Response(c, customErr.Code, customErr.Message, nil)
	}

	return helpers.Response(c, http.StatusCreated, "please check your email for code verification", nil)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	if c.Method() != http.MethodPost {
		return helpers.Response(c, http.StatusMethodNotAllowed, "Method Not Allowed", nil)
	}

	var userRequest user_account.LoginUserRequest

	// body parsing
	if err := c.BodyParser(&userRequest); err != nil {
		return helpers.Response(c, http.StatusBadRequest, "Bad Request", nil)
	}

	userByEmail, err := h.UserAccountService.FetchUserByEmail(c.Context(), userRequest.Email)

	var customErr *helpers.CustomError
	if errors.As(err, &customErr) {
		return helpers.Response(c, customErr.Code, customErr.Message, nil)
	}

	err = helpers.ComparePassword(userByEmail.Password, userRequest.Password)
	if err != nil {
		return helpers.Response(c, http.StatusUnauthorized, "Unauthorized", nil)
	}

	loginUserResponse, err := h.UserAccountService.Login(c.Context(), userByEmail)
	if errors.As(err, &customErr) {
		return helpers.Response(c, customErr.Code, customErr.Message, nil)
	}

	return helpers.Response(c, http.StatusOK, "Success", loginUserResponse)

}

//func (h *AuthHandler) GeneratePassResetCode(c *fiber.Ctx) error {
//	if c.Method() != http.MethodPost {
//		return helpers.Response(c, http.StatusMethodNotAllowed, "Method Not Allowed", nil)
//	}
//
//	var forgotPassReq user_account2.ForgotPasswordRequest
//	var customErr *helpers.CustomError
//
//	err := c.BodyParser(&forgotPassReq)
//	if err != nil {
//		return helpers.Response(c, http.StatusBadRequest, "Bad Request", nil)
//	}
//
//	userByEmail, err := h.UserAccountService.FetchUserByEmail(c.Context(), forgotPassReq.Email)
//	if errors.As(err, &customErr) {
//		return helpers.Response(c, customErr.Code, customErr.Message, nil)
//	}
//
//	// send verification mail
//	from := "andrepriyanto95@gmail.com"
//	to := []string{userByEmail.Email}
//	subject := "Reset Password"
//	mailData := &user_account2.MailData{
//		Username: userByEmail.Username,
//		Code:     helpers.CodeVerification(),
//	}
//
//	mailReq := h.UserAccountService.NewMail(to, from, subject, mailData)
//	err = h.UserAccountService.SendMail(mailReq)
//	if errors.As(err, &customErr) {
//		return helpers.Response(c, customErr.Code, customErr.Message, nil)
//	}
//
//	// store the password reset code to db
//	verificationData := models.VerificationData{
//		Email:     userByEmail.Email,
//		Code:      mailData.Code,
//		ExpiredAt: time.Now().Add(time.Minute * time.Duration(1)),
//	}
//
//	err = h.UserAccountService.StoreVerificationData(c.Context(), &verificationData)
//	if err != nil {
//		return helpers.Response(c, http.StatusInternalServerError, err.Error(), nil)
//	}
//
//	return helpers.Response(c, http.StatusOK, "Please check your mail for password reset", nil)
//}
