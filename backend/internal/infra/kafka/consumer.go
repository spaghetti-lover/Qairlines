package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
)

type MailConsumer struct {
	reader *kafka.Reader
	mailer adapters.IEmailRepository
}

func NewMailConsumer(broker, topic string, groupID string, mailer adapters.IEmailRepository) *MailConsumer {
	return &MailConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{broker},
			GroupID: groupID,
			Topic:   topic,
		}),
		mailer: mailer,
	}
}

func (c *MailConsumer) Start(ctx context.Context) error {
	for {
		m, err := c.reader.ReadMessage(ctx)
		if err != nil {
			return err
		}

		var msg dto.MailMessage
		if err := json.Unmarshal(m.Value, &msg); err != nil {
			continue
		}

		err = c.mailer.Send(ctx, msg.To, msg.Subject, msg.Body)
		if err != nil {
			log.Printf("Failed to send email to %s: %v", msg.To, err)
		}
	}
}

func (c *MailConsumer) Close() error {
	return c.reader.Close()
}
