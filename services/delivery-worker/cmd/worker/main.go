package main

import (
	"log"

	"gitverse.ru/go-portfolio/notify-platform/delivery-worker/internal/kafka"
	"gitverse.ru/go-portfolio/notify-platform/delivery-worker/internal/processor"
	"gitverse.ru/go-portfolio/notify-platform/delivery-worker/internal/sender"
)

func main() {
	log.Println("Запуск обработчика доставки")

	s := &sender.LogSender{}
	processor := processor.NewProcessor(s)
	consumer := kafka.NewConsumer(processor)
	consumer.Start()
}
