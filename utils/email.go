package utils

import (
	"net/smtp"
	"sign-lottery/pkg/constants"

	"github.com/jordan-wright/email"
)

func SendEmail(addr string, data string) error {
	e := email.NewEmail()
	e.From = constants.SmtpFromEmail
	e.To = []string{addr}
	e.Subject = "签到登陆系统用户注册验证"
	e.Text = []byte(data)
	return e.Send(constants.SmtpAddr+constants.SmtpPort, smtp.PlainAuth("", constants.SmtpFromEmail, constants.SmtpAuthKey, constants.SmtpAddr))
}
