package responses

import "github.com/TheDigitalMadness/notifications-service-go/internal/domain/entity/notification"

type NotificationsResponse struct {
	Notifications []notification.Notification `json:"notifications"`
}
