package plans

import "github.com/gin-gonic/gin"

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	plansRoutes := router.Group("/:account_id/plans")
	{
		plansRoutes.POST("/", h.Create)
		plansRoutes.GET("/:id", h.GetOne)
	}
}
