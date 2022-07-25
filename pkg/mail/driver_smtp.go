package mail

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"projmural-backend/pkg/config"

	emailPKG "github.com/jordan-wright/email"
)

// SMTP 实现 email.Driver interface
type SMTP struct{}

// Send 实现 email.Driver interface 的 Send 方法
func (s *SMTP) Send(email Email, config config.MailConfig) bool {

	e := emailPKG.NewEmail()

	e.From = fmt.Sprintf("%v <%v>", email.From.Name, email.From.Address)
	e.To = email.To
	e.Bcc = email.Bcc
	e.Cc = email.Cc
	e.Subject = email.Subject
	e.Text = email.Text
	e.HTML = email.HTML

	log.Println(e)

	err := e.SendWithTLS(
		fmt.Sprintf("%v:%v", config.Smtp.Host, config.Smtp.Port),

		smtp.PlainAuth(
			"",
			config.Smtp.Username,
			config.Smtp.Password,
			config.Smtp.Host,
		),
		&tls.Config{
			ServerName: config.Smtp.Host,
		},
	)

	if err != nil {
		log.Println(err.Error())
		return false
	}
	log.Println("发件成功")
	return true
}
