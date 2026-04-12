package plans

import "github.com/gin-gonic/gin"

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	plansRoutes := router.Group("/accounts/:account_id/plans")
	{
		plansRoutes.POST("/", h.Create)
	}
}
