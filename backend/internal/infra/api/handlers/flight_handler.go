package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/flight"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
)

type FlightHandler struct {
	createFlightUseCase flight.ICreateFlightUseCase
	getFlightUseCase    flight.IGetFlightUseCase
}

func NewFlightHandler(createFlightUseCase flight.ICreateFlightUseCase, getFlightUseCase flight.IGetFlightUseCase) *FlightHandler {
	return &FlightHandler{createFlightUseCase: createFlightUseCase, getFlightUseCase: getFlightUseCase}
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

func (h *FlightHandler) GetFlight(w http.ResponseWriter, r *http.Request) {
	// Lấy ID từ query parameter
	flightIDStr := r.URL.Query().Get("id")
	if flightIDStr == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Flight ID is required.",
		})
		return
	}

	flightID, err := strconv.ParseInt(flightIDStr, 10, 64)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid Flight ID.",
		})
		return
	}

	// Gọi use case để lấy thông tin chuyến bay
	flightDetails, err := h.getFlightUseCase.Execute(r.Context(), flightID)
	if err != nil {
		if errors.Is(err, flight.ErrFlightNotFound) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Flight not found.",
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "An unexpected error occurred. Please try again later.",
		})
		return
	}

	// Trả về phản hồi thành công
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Flight details retrieved successfully.",
		"data":    flightDetails,
	})
}
