package httpController

type GetAllNotificationsByUserDto struct {
	UserID int `json:"userId" binding:"required"`
}

type GetAdminNotificationsDto struct {
	Page  int `json:"page" binding:"required,min=1"`
	Limit int `json:"limit" binding:"required,min=1"`
}
