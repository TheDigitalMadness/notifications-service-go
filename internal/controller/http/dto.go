package httpController

import "github.com/gin-gonic/gin"

func parseQueryDto[dtoType any](c *gin.Context) (dtoType, error) {
	var dto dtoType
	if err := c.ShouldBindQuery(&dto); err != nil {
		var nilAnswer dtoType
		return nilAnswer, err
	}
	return dto, nil
}

type getAllNotificationsByUserDto struct {
	userID int
}

type getAdminNotificationsDto struct {
	page  int
	limit int
}
