package kafka

import (
	"github.com/segmentio/kafka-go"
	"gitverse.ru/go-portfolio/notify-platform/delivery-worker/internal/processor"
)

type Consumer struct {
	reader *kafka.Reader
	processor *processor.Processor
}