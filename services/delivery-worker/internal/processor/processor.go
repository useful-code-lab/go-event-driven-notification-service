package processor

import (
	"gitverse.ru/go-portfolio/notify-platform/delivery-worker/internal/sender"
)

type Processor struct {
	sender sender.Sender
}
