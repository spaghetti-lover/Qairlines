package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/flight"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
)

type FlightHandler struct {
	createFlightUseCase        flight.ICreateFlightUseCase
	getFlightUseCase           flight.IGetFlightUseCase
	updateFlightTimesUseCase   flight.IUpdateFlightTimesUseCase
	getAllFlightsUseCase       flight.IGetAllFlightsUseCase
	deleteFlightUseCase        flight.IDeleteFlightUseCase
	searchFlightsUseCase       flight.ISearchFlightsUseCase
	getSuggestedFlightsUseCase flight.IGetSuggestedFlightsUseCase
}

func NewFlightHandler(createFlightUseCase flight.ICreateFlightUseCase, getFlightUseCase flight.IGetFlightUseCase, updateFlightTimesUseCase flight.IUpdateFlightTimesUseCase, getAllFlightsUseCase flight.IGetAllFlightsUseCase, deleteFlightUseCase flight.IDeleteFlightUseCase, searchFlightsUseCase flight.ISearchFlightsUseCase, getSuggestedFlightsUseCase flight.IGetSuggestedFlightsUseCase) *FlightHandler {
	return &FlightHandler{createFlightUseCase: createFlightUseCase, getFlightUseCase: getFlightUseCase, updateFlightTimesUseCase: updateFlightTimesUseCase,
		getAllFlightsUseCase: getAllFlightsUseCase, deleteFlightUseCase: deleteFlightUseCase,
		searchFlightsUseCase:       searchFlightsUseCase,
		getSuggestedFlightsUseCase: getSuggestedFlightsUseCase,
	}
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
		if errors.Is(err, adapters.ErrFlightNotFound) {
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

func (h *FlightHandler) UpdateFlightTimes(w http.ResponseWriter, r *http.Request) {
	// Kiểm tra quyền admin
	isAdmin := r.Header.Get("admin")
	if isAdmin != "true" {
		http.Error(w, "Authentication failed. Admin privileges required.", http.StatusUnauthorized)
		return
	}

	// Lấy flightID từ query parameter
	flightIDStr := r.URL.Query().Get("id")
	if flightIDStr == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	flightID, err := strconv.ParseInt(flightIDStr, 10, 64)
	if err != nil {
		http.Error(w, `{"message":"Invalid id"}`, http.StatusBadRequest)
		return
	}

	// Parse request body
	var request dto.UpdateFlightTimesRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, `{"message":"Invalid flight data. Please check the input fields."}`, http.StatusBadRequest)
		return
	}

	// Gọi use case để cập nhật thời gian chuyến bay
	response, err := h.updateFlightTimesUseCase.Execute(r.Context(), flightID, request)
	if err != nil {
		if errors.Is(err, adapters.ErrFlightNotFound) {
			http.Error(w, `{"message":"Flight not found."}`, http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf(`{"message":"An unexpected error occurred. Please try again later.%v"}`, err), http.StatusInternalServerError)
		return
	}

	// Trả về response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Flight updated successfully.",
		"flight":  response,
	})
}

func (h *FlightHandler) GetAllFlights(w http.ResponseWriter, r *http.Request) {
	// Kiểm tra quyền admin
	isAdmin := r.Header.Get("admin")
	if isAdmin != "true" {
		http.Error(w, "Authentication failed. Admin privileges required.", http.StatusUnauthorized)
		return
	}

	// Gọi use case để lấy danh sách chuyến bay
	flights, tickets, err := h.getAllFlightsUseCase.Execute(r.Context())
	if err != nil {
		http.Error(w, "An unexpected error occurred. Please try again later, "+err.Error(), http.StatusInternalServerError)
		return
	}
	response := mappers.MapFlightsAndTicketsToResponse(flights, tickets)
	// Trả về response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Flights retrieved successfully.",
		"data":    response,
	})
}

func (h *FlightHandler) DeleteFlight(w http.ResponseWriter, r *http.Request) {
	// Kiểm tra quyền admin
	isAdmin := r.Header.Get("admin")
	if isAdmin != "true" {
		http.Error(w, "Authentication failed. Admin privileges required.", http.StatusUnauthorized)
		return
	}

	// Lấy flightID từ query parameter
	flightIDStr := r.URL.Query().Get("id")
	if flightIDStr == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	flightID, err := strconv.ParseInt(flightIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	// Gọi use case để xóa chuyến bay
	err = h.deleteFlightUseCase.Execute(r.Context(), flightID)
	if err != nil {
		if errors.Is(err, adapters.ErrFlightNotFound) {
			http.Error(w, `{"message":"Flight not found."}`, http.StatusNotFound)
			return
		}
		http.Error(w, `{"message":"An unexpected error occurred. Please try again later."}`+err.Error(), http.StatusInternalServerError)
		return
	}

	// Trả về response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Flight deleted successfully."}`))
}

func (h *FlightHandler) SearchFlights(w http.ResponseWriter, r *http.Request) {
	// Lấy query parameters
	departureCity := r.URL.Query().Get("departureCity")
	arrivalCity := r.URL.Query().Get("arrivalCity")
	flightDate := r.URL.Query().Get("flightDate")

	// Kiểm tra các query parameters
	if departureCity == "" || arrivalCity == "" || flightDate == "" {
		http.Error(w, `{"message": "Invalid query parameters. Please check departureCity, arrivalCity, and flightDate."}`, http.StatusBadRequest)
		return
	}

	// Gọi use case để tìm kiếm chuyến bay
	flights, err := h.searchFlightsUseCase.Execute(r.Context(), departureCity, arrivalCity, flightDate)
	if err != nil {
		if errors.Is(err, adapters.ErrNoFlightsFound) {
			http.Error(w, `{"message": "No flights found for the given criteria."}`, http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf(`{"message": "An unexpected error occurred. Please try again later.%v"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	// Trả về response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Flights retrieved successfully.",
		"data":    flights,
	})
}

func (h *FlightHandler) GetSuggestedFlights(w http.ResponseWriter, r *http.Request) {
	// Gọi use case để lấy danh sách chuyến bay gợi ý
	flights, err := h.getSuggestedFlightsUseCase.Execute(r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"message": "An unexpected error occurred. %v"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	// Trả về response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Suggested flights retrieved successfully.",
		"data":    flights,
	})
}
