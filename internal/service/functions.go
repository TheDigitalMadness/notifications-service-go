package service

import (
	"context"
	"fmt"

	httpController "github.com/TheDigitalMadness/notifications-service-go/internal/controller/http"
	notification_model "github.com/TheDigitalMadness/notifications-service-go/internal/models/notification"
)

// CreateNotification creates notification with any publicType
func (s *service) CreateNotification(ctx context.Context, publicType notification_model.PublicType, type_ string, message string, userID *int) {
	err := s.repo.CreateNotification(ctx, publicType, type_, message, userID)
	if err != nil {
		// TODO: прикрутить логгер
		fmt.Print(err)
	}
}

// GetAllNotificationsByUserDto returns all notifications by userID
func (s *service) GetAllNotificationsByUser(ctx context.Context, userID int) (httpController.NotificationsResponse, error) {
	notifications, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return httpController.NotificationsResponse{}, err
	}

	var ids []string
	for _, notification := range notifications {
		if !notification.IsRead {
			ids = append(ids, notification.ID)
		}
	}

	err = s.repo.SetRead(ctx, ids)
	if err != nil {
		return httpController.NotificationsResponse{}, err
	}

	return httpController.NotificationsResponse{Notifications: notifications}, nil
}

// GetAdminNotifications returns all admin notifications with pagination
func (s *service) GetAdminNotifications(ctx context.Context, page int, limit int) (httpController.NotificationsResponse, error) {
	notifications, err := s.repo.GetAdminNotifications(ctx, page, limit)
	if err != nil {
		return httpController.NotificationsResponse{}, err
	}

	var ids []string
	for _, notification := range notifications {
		if !notification.IsRead {
			ids = append(ids, notification.ID)
		}
	}

	err = s.repo.SetRead(ctx, ids)
	if err != nil {
		return httpController.NotificationsResponse{}, err
	}

	return httpController.NotificationsResponse{Notifications: notifications}, nil
}
