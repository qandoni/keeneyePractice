package core_email

import (
	"context"
	"fmt"

	"github.com/wneessen/go-mail"
)

type SMTPSender struct {
	host     string
	port     int
	username string
	password string
	from     string
}

func NewSMTPSender(config Config) *SMTPSender {
	return &SMTPSender{
		host:     config.Host,
		port:     config.Port,
		username: config.Username,
		password: config.Password,
		from:     config.From,
	}
}

func (s *SMTPSender) SendRegistrationEmail(
	ctx context.Context,
	to string,
	subject string,
	body string,
) error {

	msg := mail.NewMsg()

	if err := msg.From(s.from); err != nil {
		return fmt.Errorf("from: %w", err)
	}

	if err := msg.To(to); err != nil {
		return fmt.Errorf("to: %w", err)
	}

	msg.Subject(subject)
	msg.SetBodyString(mail.TypeTextPlain, body)

	client, err := mail.NewClient(
		s.host,
		mail.WithPort(s.port),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(s.username),
		mail.WithPassword(s.password),
	)
	if err != nil {
		return fmt.Errorf("create client: %w", err)
	}

	if err := client.DialAndSend(msg); err != nil {
		return fmt.Errorf("send email: %w", err)
	}

	return nil
}
