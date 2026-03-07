package handler

import (
	"context"

	"gitverse.ru/go-portfolio/notify-platform/notification-command-service/internal/service"
	notificationpb "gitverse.ru/go-portfolio/notify-platform/notification-command-service/proto"
)

type Handler struct {
	notificationpb.UnimplementedNotificationServiceServer
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) SendNotification(
	ctx context.Context,
	req *notificationpb.SendNotificationRequest,
) (*notificationpb.SendNotificationResponse, error) {
	id, err := h.service.SendNotification(ctx, req)

	if err != nil {
		return nil, err
	}

	return &notificationpb.SendNotificationResponse{
		Id: id,
	}, nil
}
