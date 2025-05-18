package handlers

// import (
// 	"context"
// 	"encoding/json"
// 	"net/http"
// 	"runtime"
// 	"time"

// 	"github.com/spaghetti-lover/qairlines/pkg/utils"
// )

// type Stats struct {
// 	CPUPercent string `json:"cpu_percent"`
// 	CPUCore    int    `json:"cpu_core"`
// }
// type HealthResponse struct {
// 	Status  string `json:"status"`
// 	Version string `json:"version,omitempty"`
// 	Stats   Stats  `json:"stats,omitempty"`
// }

// // Health check handler
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
