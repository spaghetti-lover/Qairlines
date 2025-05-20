package adapters

import "context"

type IEmailRepository interface {
	Send(context context.Context, to, subject, body string) error
}
