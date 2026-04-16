package subscription

import "github.com/gin-gonic/gin"

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	subscriptions := router.Group("/:account_id/subscriptions")
	{
		subscriptions.POST("/", h.Create)
	}
}
