package mailer

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/kafka"
)

type MailRepository struct {
	producer *kafka.MailProducer
}

func NewMailRepository(producer *kafka.MailProducer) *MailRepository {
	return &MailRepository{
		producer: producer,
	}
}

func (r *MailRepository) Send(ctx context.Context, to, subject, body string) error {
	emailMessage := dto.MailMessage{
		To:      to,
		Subject: subject,
		Body:    body,
	}
	return r.producer.Publish(ctx, emailMessage)
}
