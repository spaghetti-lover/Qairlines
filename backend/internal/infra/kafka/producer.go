package kafka

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
)

type MailProducer struct {
	writer *kafka.Writer
}

func NewMailProducer(broker, topic string) *MailProducer {
	return &MailProducer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(broker),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (p *MailProducer) Publish(ctx context.Context, msg dto.MailMessage) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return p.writer.WriteMessages(ctx, kafka.Message{Value: data})
}
