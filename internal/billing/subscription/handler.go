package subscription

import (
	"errors"
	"net/http"

	"github.com/EduardoMark/BillingCore/pkg/validate"
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

	var payload CreateSubscriptionRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := validate.Validate(payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	subscription, err := h.service.Create(c, accountID, &payload)
	if err != nil {
		if errors.Is(err, ErrCustomerAlreadyHasSubscription) {
			c.JSON(http.StatusConflict, gin.H{"error": "Customer already has a subscription for this plan"})
			return
		}

		zap.L().Error("Handler.Create error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create subscription"})
		return
	}

	c.JSON(http.StatusCreated, subscription)
}
