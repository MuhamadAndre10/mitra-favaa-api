package handler

import (
	"errors"
	"github.com/andrepriyanto10/favaa_mitra/internal/domain/authentication/model"
	"github.com/andrepriyanto10/favaa_mitra/internal/domain/authentication/services"
	"github.com/andrepriyanto10/favaa_mitra/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type AuthHandler struct {
	model.UserService
}

func NewAuthHandler(service *services.AuthService) AuthHandler {
	return AuthHandler{
		UserService: service,
	}

}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	if c.Method() != http.MethodPost {
		return utils.Response(c, http.StatusMethodNotAllowed, "Method Not Allowed", nil)
	}

	userRequest := new(model.UserRegisterReq)

	// body parsing
	if err := c.BodyParser(userRequest); err != nil {
		return utils.Response(c, http.StatusBadRequest, "Bad Request", nil)
	}

	res, err := h.UserService.Register(c.Context(), userRequest)

	var customErr *utils.CustomError
	if errors.As(err, &customErr) {
		return utils.Response(c, customErr.Code, customErr.Message, nil)
	}

	return utils.Response(c, http.StatusCreated, "Success", res)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	if c.Method() != http.MethodPost {
		return utils.Response(c, http.StatusMethodNotAllowed, "Method Not Allowed", nil)
	}

	var userRequest model.LoginUserRequest

	// body parsing
	if err := c.BodyParser(&userRequest); err != nil {
		return utils.Response(c, http.StatusBadRequest, "Bad Request", nil)
	}

	userByEmail, err := h.UserService.FetchUserByEmail(c.Context(), userRequest.Email)

	var customErr *utils.CustomError
	if errors.As(err, &customErr) {
		return utils.Response(c, customErr.Code, customErr.Message, nil)
	}
	
	err = utils.ComparePassword(userByEmail.Password, userRequest.Password)
	if err != nil {
		return utils.Response(c, http.StatusUnauthorized, "Unauthorized", nil)
	}

	loginUserResponse, err := h.UserService.Login(c.Context(), userByEmail)
	if errors.As(err, &customErr) {
		return utils.Response(c, customErr.Code, customErr.Message, nil)
	}

	return utils.Response(c, http.StatusOK, "Success", loginUserResponse)

}
