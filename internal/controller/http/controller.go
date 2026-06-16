package httpController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetAllNotificationsByUser(ctx *gin.Context) {
	dto, err := parseQueryDto[getAllNotificationsByUserDto](ctx)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err,
			},
		)
		return
	}

	resp, err := h.service.GetAllNotificationsByUser(ctx.Request.Context(), dto.userID)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "Internal server error",
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		resp,
	)
}

func (h *handler) GetAdminNotifications(ctx *gin.Context) {
	dto, err := parseQueryDto[getAdminNotificationsDto](ctx)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err,
			},
		)
		return
	}

	resp, err := h.service.GetAdminNotifications(ctx.Request.Context(), dto.page, dto.limit)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "Internal server error",
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		resp,
	)
}
