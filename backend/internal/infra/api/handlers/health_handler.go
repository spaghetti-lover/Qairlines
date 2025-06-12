package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases"
)

type HealthHandler struct {
	healthUseCase usecases.IHealthUseCase
}

func NewHealthHandler(healthUseCase usecases.IHealthUseCase) *HealthHandler {
	return &HealthHandler{
		healthUseCase: healthUseCase,
	}
}

func (h *HealthHandler) GetHealth(c *gin.Context) {
	health, err := h.healthUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get health status", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, health)
}
