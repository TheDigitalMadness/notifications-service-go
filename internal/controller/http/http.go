package httpController

import (
	"context"

	"github.com/TheDigitalMadness/notifications-service-go/internal/domain/responses"
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
