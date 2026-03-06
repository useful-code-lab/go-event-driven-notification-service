package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
	"gitverse.ru/go-portfolio/notify-platform/delivery-worker/internal/domain"
	"gitverse.ru/go-portfolio/notify-platform/delivery-worker/internal/processor"
)

type Consumer struct {
	reader    *kafka.Reader
	processor *processor.Processor
}

func NewConsumer(processor *processor.Processor) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "notifications.created",
		GroupID: "delivery-worker",
	})

	return &Consumer{
		reader:    reader,
		processor: processor,
	}
}

func (c *Consumer) Start() {
	for {
		msg, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Ошибка чтения Kafka: ", err)
			continue
		}

		var notification domain.Notification

		err = json.Unmarshal(msg.Value, &notification)
		if err != nil {
			log.Println("Ошибка декодирования: ", err)
			continue
		}

		c.processor.Process(notification)
	}
}
