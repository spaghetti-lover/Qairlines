package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/ticket"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
)

type TicketHandler struct {
	getTicketsByFlightIDUseCase ticket.IGetTicketsByFlightIDUseCase
	getTicketUseCase            ticket.IGetTicketUseCase
	cancelTicketUseCase         ticket.ICancelTicketUseCase
	updateSeatsUseCase          ticket.IUpdateSeatsUseCase
}

func NewTicketHandler(getTicketsByFlightIDUseCase ticket.IGetTicketsByFlightIDUseCase, getTicketUseCase ticket.IGetTicketUseCase, cancelTicketUseCase ticket.ICancelTicketUseCase, updateSeatsUseCase ticket.IUpdateSeatsUseCase) *TicketHandler {
	return &TicketHandler{
		getTicketsByFlightIDUseCase: getTicketsByFlightIDUseCase,
		getTicketUseCase:            getTicketUseCase,
		cancelTicketUseCase:         cancelTicketUseCase,
		updateSeatsUseCase:          updateSeatsUseCase,
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

func (h *TicketHandler) GetTicket(w http.ResponseWriter, r *http.Request) {
	ticketIDStr := r.URL.Query().Get("id")
	if ticketIDStr == "" {
		http.Error(w, `{"message": "id is required"}`, http.StatusBadRequest)
		return
	}
	ticketID, err := strconv.ParseInt(ticketIDStr, 10, 64)
	if err != nil {
		http.Error(w, `{"message":"Invalid id"}`, http.StatusBadRequest)
		return
	}
	ticket, err := h.getTicketUseCase.Execute(r.Context(), ticketID)
	if err != nil {
		if errors.Is(err, adapters.ErrTicketNotFound) {
			http.Error(w, "Ticket not found", http.StatusNotFound)
			return
		}
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Ticket retrieved successfully.",
		"data":    ticket,
	})
}

func (h *TicketHandler) CancelTicket(w http.ResponseWriter, r *http.Request) {

	ticketIDStr := r.URL.Query().Get("id")
	if ticketIDStr == "" {
		http.Error(w, `{"message":"id is required"}`, http.StatusBadRequest)
		return
	}

	ticketID, err := strconv.ParseInt(ticketIDStr, 10, 64)
	if err != nil {
		http.Error(w, `{"message":"Invalid id"}`, http.StatusBadRequest)
		return
	}

	ticket, err := h.cancelTicketUseCase.Execute(r.Context(), ticketID)
	if err != nil {
		if errors.Is(err, adapters.ErrTicketNotFound) {
			http.Error(w, `{"message":"Ticket not found"}`, http.StatusNotFound)
			return
		}
		if errors.Is(err, adapters.ErrTicketCannotBeCancelled) {
			http.Error(w, `{"message":"Ticket cannot be cancelled due to its current status."}`, http.StatusBadRequest)
			return
		}
		http.Error(w, `{"message":"An unexpected error occurred"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Ticket cancelled successfully.",
		"ticket":  ticket,
	})
}

func (h *TicketHandler) UpdateSeats(w http.ResponseWriter, r *http.Request) {
	var updates []dto.UpdateSeatRequest
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid seat data. Please check the input fields."+err.Error(), http.StatusBadRequest)
		return
	}

	responses, err := h.updateSeatsUseCase.Execute(r.Context(), updates)
	if err != nil {
		if errors.Is(err, adapters.ErrTicketNotFound) {
			http.Error(w, `{"message":"One or more tickets not found."}`, http.StatusNotFound)
			return
		}
		if errors.Is(err, adapters.ErrInvalidSeat) {
			http.Error(w, `{"message":"Invalid seat data. Please check the input fields."}`, http.StatusBadRequest)
			return
		}
		http.Error(w, `{"message":"An unexpected error occurred. Please try again later."}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Seats updated successfully.",
		"data":    mappers.ToUpdateSeatResponses(responses),
	})
}
