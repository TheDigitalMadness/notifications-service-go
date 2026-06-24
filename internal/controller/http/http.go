package httpController

import (
	"context"
)

type HttpService interface {
	GetAllNotificationsByUser(ctx context.Context, userID int) (NotificationsResponse, error)
	GetAdminNotifications(ctx context.Context, page int, limit int) (NotificationsResponse, error)
}

type handler struct {
	service HttpService
}

func New(service HttpService) *handler {
	return &handler{service: service}
}
