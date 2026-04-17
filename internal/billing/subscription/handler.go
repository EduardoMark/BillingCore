package subscription

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/EduardoMark/BillingCore/internal/cache"
	"github.com/EduardoMark/BillingCore/pkg/validate"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	service *Service
	cache   cache.Cache
}

func NewHandler(service *Service, cache cache.Cache) *Handler {
	return &Handler{
		service: service,
		cache:   cache,
	}
}

func (h *Handler) Create(c *gin.Context) {
	zap.L().Info("Handler.Create running")
	accountID := c.Param("account_id")
	idempotencyKey := c.GetHeader("Idempotency-Key")
	if idempotencyKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Idempotency-Key is required"})
		return
	}

	var cached IdempotencyValue
	ok, err := h.cache.SetNX(c, idempotencyKey, IdempotencyValue{Status: "processing"}, time.Hour*1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !ok {
		fmt.Printf("Idempotency key %s is already being processed\n", idempotencyKey)
		found, _ := h.cache.Get(c, idempotencyKey, &cached)
		if found && cached.Status == "processing" {
			c.JSON(http.StatusAccepted, gin.H{"message": "Subscription creation is being processed"})
			return
		}

		c.JSON(http.StatusOK, cached.Data)
		return
	}
	fmt.Printf("Idempotency key %s is set to processing\n", idempotencyKey)
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

	if err := h.cache.Set(c, idempotencyKey, IdempotencyValue{Status: "completed", Data: subscription}, time.Hour*1); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, subscription)
}
