package notification_model

import "time"

type Notification struct {
	ID         string
	PublicType PublicType
	UserID     *int // Эквивалент int | null
	Type       string
	Message    string
	IsRead     bool
	CreatedAt  time.Time
}
