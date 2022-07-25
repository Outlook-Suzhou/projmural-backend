package mail

import "projmural-backend/pkg/config"

type Driver interface {
	Send(email Email, config config.MailConfig) bool
}
