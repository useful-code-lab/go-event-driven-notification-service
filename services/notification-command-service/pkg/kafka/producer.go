package kafka

import "github.com/segmentio/kafka-go"

type Producer struct {
	writer *kafka.Writer
}