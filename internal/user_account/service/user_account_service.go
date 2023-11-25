package service

import (
	"context"
	"github.com/andrepriyanto10/favaa_mitra/internal/helpers"
	"github.com/andrepriyanto10/favaa_mitra/internal/tool/mail"
	"github.com/andrepriyanto10/favaa_mitra/internal/tool/token"
	"github.com/andrepriyanto10/favaa_mitra/internal/user_account"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/andrepriyanto10/favaa_mitra/internal/models"
	"github.com/spf13/viper"
)

type Service struct {
	user    user_account.UserAccountRepository
	timeout time.Duration
	env     *viper.Viper
}

func NewAuthService(user user_account.UserAccountRepository, env *viper.Viper) *Service {
	return &Service{
		user:    user,
		timeout: time.Duration(4) * time.Second,
		env:     env,
	}
}

func (u *Service) VerificationCode(ctx context.Context, user *user_account.UserRegisterReq) error {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	if user.Username == "" || user.Email == "" || user.Password == "" {
		return &helpers.CustomError{Code: http.StatusBadRequest, Message: "Bad Request"}
	}

	if len(user.Username) < 5 || len(user.Username) > 32 {
		return &helpers.CustomError{Code: http.StatusBadRequest, Message: "Username must be between 5 and 32 characters"}
	}

	// validation email and sending error response
	if !strings.Contains(strings.TrimSpace(user.Email), "@") {
		return &helpers.CustomError{Code: http.StatusBadRequest, Message: "Invalid email"}
	}
	// validation password and sending error response
	if len(user.Password) < 8 || len(user.Password) > 32 {
		return &helpers.CustomError{Code: http.StatusBadRequest, Message: "Password must be between 8 and 32 characters"}
	}

	// setup email data
	from := "andrepriyanto95@gmail.com"
	to := []string{user.Email}
	subject := "Email verification"
	mailData := &mail.MailData{
		Username: user.Username,
		Code:     helpers.CodeVerification(),
	}

	mailArg := mail.Mail{
		From:    from,
		To:      to,
		Subject: subject,
		Data:    mailData,
		Env:     u.env,
	}

	newMail := mail.NewMail(mailArg)

	// get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		return &helpers.CustomError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	err = newMail.ParseTemplate(dir+"/public/template/reset_email_template_simple.html", mailData)
	if err != nil {
		return &helpers.CustomError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	c := make(chan error, 1)

	go func() {
		c <- newMail.SendEmail()
	}()

	select {
	case err := <-c:
		if err != nil {
			return &helpers.CustomError{Code: http.StatusInternalServerError, Message: err.Error()}
		}
	case <-ctx.Done():
		return &helpers.CustomError{Code: http.StatusInternalServerError, Message: ctx.Err().Error()}
	}

	verificationData := models.VerificationData{
		Email:     user.Email,
		Code:      mailData.Code,
		ExpiredAt: time.Now().Add(time.Minute * time.Duration(1)),
	}

	err = u.user.StoreVerificationData(ctx, &verificationData)
	if err != nil {
		return &helpers.CustomError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return nil

}

func (u *Service) FetchUserByEmail(ctx context.Context, email string) (*models.UserAccounts, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	if !helpers.Valid(email) {
		return nil, &helpers.CustomError{Code: http.StatusBadRequest, Message: "Bad Request"}
	}

	if email == "" || !strings.Contains(strings.TrimSpace(email), "@") {
		return nil, &helpers.CustomError{Code: http.StatusBadRequest, Message: "Bad Request"}
	}

	user, err := u.user.FetchUserByEmail(ctx, email)
	if err != nil {
		return nil, &helpers.CustomError{Code: http.StatusInternalServerError, Message: "User not found"}
	}

	return user, nil
}

func (u *Service) Login(ctx context.Context, user *models.UserAccounts) (*user_account.LoginUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	tokenExpired, _ := time.ParseDuration(u.env.GetString("JWT_TTL"))

	// token token
	tkn, err := token.GenerateToken(tokenExpired, user.ID, u.env.GetString("JWT_SECRET"))
	if err != nil {
		return nil, &helpers.CustomError{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}

	loginRes := &user_account.LoginUserResponse{
		Token: tkn,
	}

	return loginRes, nil

}

func (u *Service) StoreVerificationData(ctx context.Context, data *models.VerificationData) error {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	err := u.user.StoreVerificationData(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
