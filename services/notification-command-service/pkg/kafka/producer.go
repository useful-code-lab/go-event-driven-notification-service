package kafka

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
	"gitverse.ru/go-portfolio/notify-platform/notification-command-service/internal/domain"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(brokers []string, topic string) *Producer{
	return &Producer{
		writer: &kafka.Writer{
			Addr: kafka.TCP(brokers...),
			Topic: topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (p *Producer) Publish(
	notification domain.Notification,
) error {
	data, _ := json.Marshal(notification)

	msg := kafka.Message{
		Value: data,
	}

	return p.writer.WriteMessages(
		context.Background(),
		msg,
	)
}
