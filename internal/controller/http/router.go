package httpController

import "github.com/gin-gonic/gin"

type HttpHandler interface {
	GetAllNotificationsByUser(ctx *gin.Context)
	GetAdminNotifications(ctx *gin.Context)
}

func NewRouter(handler HttpHandler) *gin.Engine {
	router := gin.Default()

	router.GET("by-user", handler.GetAllNotificationsByUser)
	router.GET("admin", handler.GetAdminNotifications)

	return router
}
