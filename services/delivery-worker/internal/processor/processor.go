package processor

import (
	"log"

	"gitverse.ru/go-portfolio/notify-platform/delivery-worker/internal/domain"
	"gitverse.ru/go-portfolio/notify-platform/delivery-worker/internal/sender"
)

type Processor struct {
	sender sender.Sender
}

func NewProcessor(sender sender.Sender) *Processor {
	return &Processor{
		sender: sender,
	}
}

func (p *Processor) Process(n domain.Notification) {
	log.Println("Уведомление обработки: ", n.ID)

	err := p.sender.Send(n)

	if err != nil {
		log.Println("Ошибка отправки: ", err)
	}
}
