package account

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
	return &Handler{
		service: service,
	}
}

func (h *Handler) Create(c *gin.Context) {
	zap.L().Info("Handler.Create running")
	var payload CreateAccountPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		zap.L().Error("Failed to bind JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := validate.Validate(payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	account, err := h.service.Create(c, &payload)
	if err != nil {
		zap.L().Error("Failed to create account", zap.Error(err))

		if errors.Is(err, ErrEmailAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create account"})
		return
	}

	c.JSON(http.StatusCreated, account)
}

func (h *Handler) GetOne(c *gin.Context) {
	zap.L().Info("Handler.GetOne running")
	id := c.Param("id")

	account, err := h.service.GetByID(c, id)
	if err != nil {
		if errors.Is(err, ErrAccountNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
			return
		}

		zap.L().Error("Failed to get account", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get account"})
		return
	}

	c.JSON(http.StatusOK, account)
}
