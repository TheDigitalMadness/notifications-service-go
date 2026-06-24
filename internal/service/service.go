package service

import (
	"context"

	notification_model "github.com/TheDigitalMadness/notifications-service-go/internal/models/notification"
)

type Repositrory interface {
	CreateNotification(ctx context.Context, publicType notification_model.PublicType, type_ string, message string, userID *int) error
	GetByUserID(ctx context.Context, userID int) ([]notification_model.Notification, error)
	GetAdminNotifications(ctx context.Context, page int, limit int) ([]notification_model.Notification, error)
	SetRead(ctx context.Context, ids []string) error
}

type service struct {
	repo Repositrory
}

func New(repo Repositrory) *service {
	return &service{repo: repo}
}
