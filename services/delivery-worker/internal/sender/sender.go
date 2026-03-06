package sender

import "gitverse.ru/go-portfolio/notify-platform/delivery-worker/internal/domain"

type Sender interface {
	Send(n domain.Notification) error
}
