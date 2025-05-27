package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/ticket"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
)

type TicketHandler struct {
	getTicketsByFlightIDUseCase ticket.IGetTicketsByFlightIDUseCase
}

func NewTicketHandler(getTicketsByFlightIDUseCase ticket.IGetTicketsByFlightIDUseCase) *TicketHandler {
	return &TicketHandler{
		getTicketsByFlightIDUseCase: getTicketsByFlightIDUseCase,
	}
}

func (h *TicketHandler) GetTicketsByFlightID(w http.ResponseWriter, r *http.Request) {
	// Kiểm tra quyền admin
	isAdmin := r.Header.Get("admin")
	if isAdmin != "true" {
		http.Error(w, `{"message": "Authentication failed. Admin privileges required."}`, http.StatusUnauthorized)
		return
	}

	// Lấy flightID từ query
	flightIDStr := r.URL.Query().Get("flightId")

	if flightIDStr == "" {
		http.Error(w, `{"message": "Flight ID is required."}`, http.StatusBadRequest)
		return
	}

	flightID, err := strconv.ParseInt(flightIDStr, 10, 64)
	if err != nil {
		http.Error(w, `{"message": "Invalid Flight ID."}`, http.StatusBadRequest)
		return
	}

	// Gọi use case
	tickets, err := h.getTicketsByFlightIDUseCase.Execute(r.Context(), flightID)
	if err != nil {
		if err == adapters.ErrFlightNotFound {
			http.Error(w, `{"message": "Flight not found."}`, http.StatusNotFound)
			return
		}
		http.Error(w, `{"message": "An unexpected error occurred. Please try again later."}`, http.StatusInternalServerError)
		return
	}

	// Tạo response
	response := mappers.ToGetTicketsByFlightIDResponse(tickets)
	// Trả về response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
