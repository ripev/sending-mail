package sending_mail

import (
	"fmt"
	"net/smtp"
)

type (
	request struct {
		from    string
		to      []string
		subject string
		body    string
		config  *Config
		message string
	}

	Config struct {
		Server string
		Port   int
		Email  string
	}
)

const (
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

func NewRequest(to []string, subject, message string, config *Config) *request {
	return &request{
		to:      to,
		subject: subject,
		message: message,
		config:  config,
	}
}

func (r *request) Send() error {
	r.body = getTemplate(cssBootstrap, r.subject, r.message)

	body := "To: " + r.to[0] + "\r\nSubject: " + r.subject + "\r\n" + MIME + "\r\n" + r.body
	SMTP := fmt.Sprintf("%s:%d", r.config.Server, r.config.Port)

	return smtp.SendMail(SMTP, nil, r.config.Email, r.to, []byte(body))
}
