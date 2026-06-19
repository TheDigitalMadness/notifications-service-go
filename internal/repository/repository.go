package repository

import (
	"context"
	"fmt"

	"github.com/TheDigitalMadness/notifications-service-go/internal/domain/entity/notification"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *repository {
	return &repository{
		pool: pool,
	}
}

func (r *repository) CreateNotification(ctx context.Context, publicType notification.PublicType, type_ string, message string, userID *int) error {
	op := "CreateNotification"

	_, err := r.pool.Exec(ctx, createNotificationQuery, publicType, userID, type_, message)
	if err != nil {
		return fmt.Errorf("Postgres query: %s: %w", op, err)
	}

	return nil
}

func (r *repository) GetByUserID(ctx context.Context, userID int) ([]notification.Notification, error) {
	op := "GetByUserID"

	rows, err := r.pool.Query(ctx, getByUserID, userID)
	if err != nil {
		return nil, fmt.Errorf("Postgres query: %s: %w", op, err)
	}

	var notifications []notification.Notification

	for rows.Next() {
		var notification notification.Notification

		err := rows.Scan(&notification)
		if err != nil {
			return nil, fmt.Errorf("Postgres query: Rows scanning: %s: %w", op, err)
		}

		notifications = append(notifications, notification)
	}

	return notifications, nil
}

func (r *repository) GetAdminNotifications(ctx context.Context, page int, limit int) ([]notification.Notification, error) {
	op := "GetAdminNotifications"

	rows, err := r.pool.Query(ctx, getAdminNotifications, notification.PublicTypeAdmin, (page-1)*limit, limit)
	if err != nil {
		return nil, fmt.Errorf("Postgres query: %s: %w", op, err)
	}

	var notifications []notification.Notification

	for rows.Next() {
		var notification notification.Notification

		err := rows.Scan(&notification)
		if err != nil {
			return nil, fmt.Errorf("Postgres query: Rows scanning: %s: %w", op, err)
		}

		notifications = append(notifications, notification)
	}

	return notifications, nil
}

func (r *repository) SetRead(ctx context.Context, ids []string) error {
	op := "SetRead"

	_, err := r.pool.Exec(ctx, setReadQuery, ids)
	if err != nil {
		return fmt.Errorf("Postgres query: %s: %w", op, err)
	}

	return nil
}
