package service

import (
	"context"

	"github.com/TheDigitalMadness/notifications-service-go/internal/domain/entity/notification"
)

type Repositrory interface {
	CreateNotification(ctx context.Context, publicType notification.PublicType, type_ string, message string, userID *int) error
	GetByUserID(ctx context.Context, userID int) ([]notification.Notification, error)
	GetAdminNotifications(ctx context.Context, page int, limit int) ([]notification.Notification, error)
	SetRead(ctx context.Context, ids []string) error
}

type service struct {
	repo Repositrory
}

func New(repo Repositrory) *service {
	return &service{repo: repo}
}
