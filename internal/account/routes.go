package account

import "github.com/gin-gonic/gin"

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	accountRoutes := router.Group("/accounts")
	{
		accountRoutes.POST("/", h.Create)
	}
}
