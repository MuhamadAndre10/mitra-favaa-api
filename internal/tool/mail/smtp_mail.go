package mail

import (
	"bytes"
	"github.com/spf13/viper"
	"net/smtp"
	"text/template"
)

type MailData struct {
	Username string
	Code     string
}

type Mail struct {
	From    string
	To      []string
	Subject string
	Body    string
	Env     *viper.Viper
	Data    *MailData
}

func NewMail(mailArg Mail) *Mail {
	return &Mail{
		To:      mailArg.To,
		Subject: mailArg.Subject,
		Body:    mailArg.Body,
		Env:     mailArg.Env,
		Data:    mailArg.Data,
	}
}

func (r *Mail) SendEmail() error {

	auth := smtp.PlainAuth(
		"",
		"andrepriyanto95@gmail.com",
		r.Env.GetString("MAIL_PASS"),
		r.Env.GetString("MAIL_HOST"))

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.Subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.Body)
	addr := r.Env.GetString("MAIL_HOST") + ":" + r.Env.GetString("MAIL_PORT")

	err := smtp.SendMail(addr, auth, r.From, r.To, msg)
	if err != nil {
		return err
	}

	return nil
}

func (r *Mail) ParseTemplate(tFileName string, data any) error {
	t, err := template.ParseFiles(tFileName)
	if err != nil {
		return err
	}

	//buf := new(strings.Builder)
	buff := new(bytes.Buffer)

	err = t.Execute(buff, data)
	if err != nil {
		return err
	}

	r.Body = buff.String()
	return nil

}
