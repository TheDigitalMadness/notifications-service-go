package httpController

import (
	"context"

	responses "github.com/TheDigitalMadness/notifications-service-go/internal/domain/response"
)

type HttpService interface {
	GetAllNotificationsByUser(ctx context.Context, userID int) (responses.NotificationsResponse, error)
	GetAdminNotifications(ctx context.Context, page int, limit int) (responses.NotificationsResponse, error)
}

type handler struct {
	service HttpService
}

func New(service HttpService) *handler {
	return &handler{service: service}
}
