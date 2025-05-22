package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/spaghetti-lover/qairlines/internal/domain/usecases"
	"github.com/spaghetti-lover/qairlines/pkg/utils"
)

type HealthHandler struct {
	healthUseCase usecases.IHealthUseCase
}

func NewHealthHandler(healthUseCase usecases.IHealthUseCase) *HealthHandler {
	return &HealthHandler{
		healthUseCase: healthUseCase,
	}
}

func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	health, err := h.healthUseCase.Execute()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "failed to get health status", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(health); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "failed to encode response", err)
		return
	}
}
