package httpController

import notification_model "github.com/TheDigitalMadness/notifications-service-go/internal/models/notification"

type NotificationsResponse struct {
	Notifications []notification_model.Notification `json:"notifications"`
}
