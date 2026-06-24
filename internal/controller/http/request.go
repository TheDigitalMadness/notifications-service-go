package httpController

type getAllNotificationsByUserDto struct {
	userID int
}

type getAdminNotificationsDto struct {
	page  int
	limit int
}
