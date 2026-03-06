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

func (p *Producer) Publish(
	notification domain.Notification,
) error {
	data, _ := json.Marshal(notification)

	msg := kafka.Message{
		Topic: "notofications",
		Value: data,
	}

	return p.writer.WriteMessages(
		context.Background(),
		msg,
	)
}
