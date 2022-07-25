package mail

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"projmural-backend/pkg/config"
	"strings"

	emailPKG "github.com/jordan-wright/email"
)

type loginAuth struct {
	username, password string
}

// loginAuth returns an Auth that implements the LOGIN authentication
// mechanism as defined in RFC 4616.
func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", nil, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	command := string(fromServer)
	command = strings.TrimSpace(command)
	command = strings.TrimSuffix(command, ":")
	command = strings.ToLower(command)

	if more {
		if command == "username" {
			return []byte(fmt.Sprintf("%s", a.username)), nil
		} else if command == "password" {
			return []byte(fmt.Sprintf("%s", a.password)), nil
		} else {
			// We've already sent everything.
			return nil, fmt.Errorf("unexpected server challenge: %s", command)
		}
	}
	return nil, nil
}

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

	err := e.SendWithStartTLS(
		fmt.Sprintf("%v:%v", config.Smtp.Host, config.Smtp.Port),
		LoginAuth(config.Smtp.Username, config.Smtp.Password),
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
