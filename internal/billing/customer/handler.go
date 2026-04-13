package customer

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
	var payload CreateCustomerPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := Validate(payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	customer, err := h.service.Create(c, accountID, &payload)
	if err != nil {
		zap.L().Error("Handler.Create error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (h *Handler) GetByID(c *gin.Context) {
	zap.L().Info("Handler.GetByID running")

	id := c.Param("id")

	customer, err := h.service.GetByID(c, id)
	if err != nil {
		zap.L().Error("Handler.GetByID error", zap.Error(err))

		if errors.Is(err, ErrCustomerNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get customer"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (h *Handler) GetByExternalID(c *gin.Context) {
	zap.L().Info("Handler.GetByExternalID running")

	externalID := c.Param("external_id")

	customer, err := h.service.GetByExternalID(c, externalID)
	if err != nil {
		zap.L().Error("Handler.GetByExternalID error", zap.Error(err))

		if errors.Is(err, ErrCustomerNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get customer"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (h *Handler) GetAllByAccountID(c *gin.Context) {
	zap.L().Info("Handler.GetAllByAccountID running")

	accountID := c.Param("account_id")

	customers, err := h.service.GetAllByAccountID(c, accountID)
	if err != nil {
		zap.L().Error("Handler.GetAllByAccountID error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get customers"})
		return
	}

	c.JSON(http.StatusOK, customers)
}

func (h *Handler) Update(c *gin.Context) {
	zap.L().Info("Handler.Update running")

	id := c.Param("id")
	var payload UpdateCustomerPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := Validate(payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	customer, err := h.service.Update(c, id, payload)
	if err != nil {
		zap.L().Error("Handler.Update error", zap.Error(err))

		if errors.Is(err, ErrCustomerNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update customer"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (h *Handler) Delete(c *gin.Context) {
	zap.L().Info("Handler.Delete running")

	id := c.Param("id")

	err := h.service.Delete(c, id)
	if err != nil {
		zap.L().Error("Handler.Delete error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete customer"})
		return
	}

	c.Status(http.StatusNoContent)
}
