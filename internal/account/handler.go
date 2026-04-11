package account

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
	return &Handler{
		service: service,
	}
}

func (h *Handler) Create(ctx *gin.Context) {
	zap.L().Info("Handler.Create running")
	var payload CreateAccountPayload

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		zap.L().Error("Failed to bind JSON", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := payload.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	account, err := h.service.Create(&payload)
	if err != nil {
		zap.L().Error("Failed to create account", zap.Error(err))

		if errors.Is(err, ErrEmailAlreadyExists) {
			ctx.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create account"})
		return
	}

	ctx.JSON(http.StatusCreated, account)
}
