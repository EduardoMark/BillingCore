package customer

import "github.com/gin-gonic/gin"

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	customers := router.Group("/:account_id/customers")
	{
		customers.POST("/", h.Create)
		customers.GET("/", h.GetAllByAccountID)
		customers.GET("/:id", h.GetByID)
		customers.GET("/external/:external_id", h.GetByExternalID)
		customers.PUT("/:id", h.Update)
		customers.DELETE("/:id", h.Delete)
	}
}
