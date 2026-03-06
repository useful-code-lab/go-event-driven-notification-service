package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gitverse.ru/go-portfolio/notify-platform/notification-command-service/internal/domain"
	"gitverse.ru/go-portfolio/notify-platform/notification-command-service/internal/repository"
	"gitverse.ru/go-portfolio/notify-platform/notification-command-service/pkg/kafka"
	notificationpb "gitverse.ru/go-portfolio/notify-platform/notification-command-service/proto"
)

type Service struct {
	repo  repository.Repository
	kafka kafka.Producer
}

func (s *Service) SendNotification(
	ctx context.Context,
	req *notificationpb.SendNotificationRequest,
) (string, error) {
	id := uuid.New().String()

	notification := domain.Notification{
		ID:       id,
		UserID:   req.UserId,
		Type:     req.Type,
		Message:  req.Message,
		Status:   "pending",
		CreateAt: time.Now(),
	}

	err := s.repo.Save(ctx, notification)
	if err != nil {
		return "", err
	}

	err = s.kafka.Publish(notification)
	if err != nil {
		return "", err
	}

	return id, nil
}
