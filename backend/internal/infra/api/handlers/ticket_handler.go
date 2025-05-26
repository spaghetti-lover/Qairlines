package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/ticket"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
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
	fmt.Print(flightIDStr)
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
		if err == ticket.ErrFlightNotFound {
			http.Error(w, `{"message": "Flight not found."}`, http.StatusNotFound)
			return
		}
		http.Error(w, `{"message": "An unexpected error occurred. Please try again later."}`, http.StatusInternalServerError)
		return
	}

	// Tạo response
	response := dto.GetTicketsResponse{
		Message: "Tickets retrieved successfully.",
		Data:    mappers.TicketsEntitiesToResponse(tickets),
	}

	// Trả về response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
