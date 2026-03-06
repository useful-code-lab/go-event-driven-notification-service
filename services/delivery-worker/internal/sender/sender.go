package sender

import (
	"log"

	"gitverse.ru/go-portfolio/notify-platform/delivery-worker/internal/domain"
)

type Sender interface {
	Send(n domain.Notification) error
}

type LogSender struct{}

func (s *LogSender) Send(n domain.Notification) error {
	log.Printf("Отправка уведомления %s пользователю: %s ", n.Message, n.UserID)

	return nil
}
