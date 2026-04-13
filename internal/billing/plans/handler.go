package plans

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(c *gin.Context) {
	zap.L().Info("Handler.Create running")
	accountID := c.Param("account_id")

	var payload CreatePlanPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := payload.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err})
		return
	}

	plan, err := h.service.Create(c, accountID, &payload)
	if err != nil {
		zap.L().Error("Handler.Create error", zap.Error(err))

		if errors.Is(err, ErrInvalidPrice) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Price must be greater than zero"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, plan)
}

func (h *Handler) GetOne(c *gin.Context) {
	zap.L().Info("Handler.GetOne running")
	id := c.Param("id")

	plan, err := h.service.GetOne(c, id)
	if err != nil {
		if errors.Is(err, ErrPlanNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Plan not found"})
			return
		}

		zap.L().Error("Handler.GetOne error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, plan)
}
