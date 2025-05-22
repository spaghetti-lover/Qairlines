package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/airplane_model"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
)

// AirplaneModelHandler handles the creation of airplane models.
type AirplaneModelHandler struct {
	airplaneModelCreateUseCase airplane_model.IAirplaneModelCreateUseCase
}

// NewAirplaneModelCreateHandler creates a new AirplaneModelCreateHandler.
func NewAirplaneModelCreateHandler(airplaneModelCreateUseCase airplane_model.IAirplaneModelCreateUseCase) *AirplaneModelHandler {
	return &AirplaneModelHandler{
		airplaneModelCreateUseCase: airplaneModelCreateUseCase,
	}
}

// ServeHTTP handles the HTTP request for creating an airplane model.
func (h *AirplaneModelHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req dto.AirplaneModelCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	input := mappers.AirplaneModelCreateRequestToInput(req)
	output, err := h.airplaneModelCreateUseCase.Execute(r.Context(), input)
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
