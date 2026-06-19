package service

import (
	"context"
	"fmt"

	"github.com/TheDigitalMadness/notifications-service-go/internal/domain/entity/notification"
	responses "github.com/TheDigitalMadness/notifications-service-go/internal/domain/response"
)

func (s *service) CreateNotification(ctx context.Context, publicType notification.PublicType, type_ string, message string, userID *int) {
	err := s.repo.CreateNotification(ctx, publicType, type_, message, userID)
	if err != nil {
		// TODO: прикрутить логгер
		fmt.Print(err)
	}
}

func (s *service) GetAllNotificationsByUserDto(ctx context.Context, userID int) (responses.NotificationsResponse, error) {
	notifications, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return responses.NotificationsResponse{}, err
	}

	var ids []string
	for _, notification := range notifications {
		if !notification.IsRead {
			ids = append(ids, notification.ID)
		}
	}

	err = s.repo.SetRead(ctx, ids)
	if err != nil {
		return responses.NotificationsResponse{}, err
	}

	return responses.NotificationsResponse{Notifications: notifications}, nil
}

func (s *service) GetAdminNotifications(ctx context.Context, page int, limit int) (responses.NotificationsResponse, error) {
	notifications, err := s.repo.GetAdminNotifications(ctx, page, limit)
	if err != nil {
		return responses.NotificationsResponse{}, err
	}

	var ids []string
	for _, notification := range notifications {
		if !notification.IsRead {
			ids = append(ids, notification.ID)
		}
	}

	err = s.repo.SetRead(ctx, ids)
	if err != nil {
		return responses.NotificationsResponse{}, err
	}

	return responses.NotificationsResponse{Notifications: notifications}, nil
}
