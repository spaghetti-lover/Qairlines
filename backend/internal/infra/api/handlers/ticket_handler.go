package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/ticket"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
)

type TicketHandler struct {
	getTicketsByFlightIDUseCase ticket.IGetTicketsByFlightIDUseCase
	cancelTicketUseCase         ticket.ICancelTicketUseCase
	getTicketUseCase            ticket.IGetTicketUseCase
}

func NewTicketHandler(getTicketsByFlightIDUseCase ticket.IGetTicketsByFlightIDUseCase, cancelTicketUseCase ticket.ICancelTicketUseCase, getTicketUseCase ticket.IGetTicketUseCase) *TicketHandler {
	return &TicketHandler{
		getTicketsByFlightIDUseCase: getTicketsByFlightIDUseCase,
		cancelTicketUseCase:         cancelTicketUseCase,
		getTicketUseCase:            getTicketUseCase,
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
	response := dto.GetTicketsResponse{
		Message: "Tickets retrieved successfully.",
		Data:    mappers.TicketsEntitiesToResponse(tickets),
	}

	// Trả về response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *TicketHandler) CancelTicket(w http.ResponseWriter, r *http.Request) {
	// Kiểm tra quyền admin
	isAdmin := r.Header.Get("admin")
	if isAdmin != "true" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Authentication failed. Admin privileges required.",
		})
		return
	}

	// Lấy ticketID từ query
	ticketIDStr := r.URL.Query().Get("id")
	if ticketIDStr == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Ticket ID is required.",
		})
		return
	}

	ticketID, err := strconv.ParseInt(ticketIDStr, 10, 64)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid Ticket ID.",
		})
		return
	}

	// Gọi use case
	ticket, err := h.cancelTicketUseCase.Execute(r.Context(), ticketID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")

		if errors.Is(err, adapters.ErrTicketNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Ticket not found.",
			})
			return
		} else if errors.Is(err, adapters.ErrTicketCannotBeCancelled) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Ticket cannot be cancelled due to its current status.",
			})
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "An unexpected error occurred. Please try again later" + fmt.Sprintf(": %v", err),
		})
		return
	}

	// Tạo response
	response := mappers.ToCancelTicketResponse(ticket)

	// Trả về response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *TicketHandler) GetTicket(w http.ResponseWriter, r *http.Request) {
	// Lấy ID từ query parameter
	ticketIDStr := r.URL.Query().Get("id")
	if ticketIDStr == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Ticket ID is required.",
		})
		return
	}

	ticketID, err := strconv.ParseInt(ticketIDStr, 10, 64)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid Ticket ID.",
		})
		return
	}

	// Gọi use case để lấy thông tin vé
	ticketDetails, err := h.getTicketUseCase.Execute(r.Context(), ticketID)
	if err != nil {
		if errors.Is(err, adapters.ErrTicketNotFound) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Ticket not found.",
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
		"message": "Ticket retrieved successfully.",
		"data":    ticketDetails,
	})
}
