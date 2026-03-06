package service

import (
	"gitverse.ru/go-portfolio/notify-platform/notification-command-service/internal/repository"
	"gitverse.ru/go-portfolio/notify-platform/notification-command-service/pkg/kafka"
)

type Service struct {
	repo  repository.Repository
	kafka kafka.Producer
}
