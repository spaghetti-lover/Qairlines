package entities

import (
	"errors"
)

type Mail struct {
	To      string
	Subject string
	Body    string
}

func NewEmail(to string, subject string, body string) (*Mail, error) {
	if to == "" {
		return nil, errors.New("username is required")
	}
	if subject == "" {
		return nil, errors.New("password is required")
	}
	if body == "" {
		return nil, errors.New("body is required")
	}
	return &Mail{
		To:      to,
		Subject: subject,
		Body:    body,
	}, nil
}
