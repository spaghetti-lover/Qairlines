package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/flight"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
)

type FlightHandler struct {
	createFlightUseCase flight.ICreateFlightUseCase
}

func NewFlightHandler(createFlightUseCase flight.ICreateFlightUseCase) *FlightHandler {
	return &FlightHandler{createFlightUseCase: createFlightUseCase}
}

func (h *FlightHandler) CreateFlight(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateFlightRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf(`{"message": "Invalid request body, %v"}`, err), http.StatusBadRequest)
		return
	}

	flightEntity := mappers.CreateFlightRequestToEntity(req)
	createdFlight, err := h.createFlightUseCase.Execute(r.Context(), flightEntity)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"message": "Failed to create flight, %v"}`, err), http.StatusInternalServerError)
		return
	}

	response := mappers.CreateFlightEntityToResponse(createdFlight)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
