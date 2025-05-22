package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/airplane_model"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
	"github.com/spaghetti-lover/qairlines/pkg/utils"
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
		utils.WriteError(w, http.StatusBadRequest, "failed to decode request body", err)
		return
	}

	input := mappers.AirplaneModelCreateRequestToInput(req)
	output, err := h.airplaneModelCreateUseCase.Execute(r.Context(), input)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "failed to create airplane model", err)
		return
	}

	response := mappers.AirplaneModelCreateOutputToResponse(output)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "failed to encode response", err)
		return
	}
}
