package mailer

import (
	"context"
	"fmt"
	"net/smtp"
)

type SMTPMailer struct {
	From     string
	Password string
	Host     string
	Port     string
}

func NewSMTPMailer(from, password, host, port string) *SMTPMailer {
	return &SMTPMailer{
		From:     from,
		Password: password,
		Host:     host,
		Port:     port,
	}
}

func (s *SMTPMailer) Send(ctx context.Context, to, subject, body string) error {
	addr := fmt.Sprintf("%s:%s", s.Host, s.Port)
	auth := smtp.PlainAuth("", s.From, s.Password, s.Host)
	msg := []byte("Subject: " + subject + "\r\n\r\n" + body)

	return smtp.SendMail(addr, auth, s.From, []string{to}, msg)
}
