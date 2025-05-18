package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/spaghetti-lover/qairlines/internal/domain/usecases"
)

// Health check handler
// func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 	defer cancel()

// 	monitor := &utils.CPUMonitor{}
// 	cpuPercent, err := monitor.Check(ctx)
// 	if err != nil {
// 		errorResponse(w, err, http.StatusInternalServerError)
// 		return
// 	}

// 	resp := HealthResponse{
// 		Status:  "OK",
// 		Version: "1.0.0",
// 		Stats: Stats{
// 			CPUPercent: cpuPercent,
// 			CPUCore:    runtime.NumCPU(),
// 		},
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	if err := json.NewEncoder(w).Encode(resp); err != nil {
// 		errorResponse(w, err, http.StatusInternalServerError)
// 		return
// 	}
// }

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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(health); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
