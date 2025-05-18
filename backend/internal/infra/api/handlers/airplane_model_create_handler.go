package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/spaghetti-lover/qairlines/internal/domain/usecases"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
)

// AirplaneModelCreateHandler handles the creation of airplane models.
type AirplaneModelCreateHandler struct {
	airplaneModelCreateUseCase usecases.IAirplaneModelCreateUseCase
}

// NewAirplaneModelCreateHandler creates a new AirplaneModelCreateHandler.
func NewAirplaneModelCreateHandler(airplaneModelCreateUseCase usecases.IAirplaneModelCreateUseCase) *AirplaneModelCreateHandler {
	return &AirplaneModelCreateHandler{
		airplaneModelCreateUseCase: airplaneModelCreateUseCase,
	}
}

// ServeHTTP handles the HTTP request for creating an airplane model.
func (h *AirplaneModelCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req dto.AirplaneModelCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	input := mappers.AirplaneModelCreateRequestToInput(req)
	output, err := h.airplaneModelCreateUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := mappers.AirplaneModelCreateOutputToResponse(output)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
