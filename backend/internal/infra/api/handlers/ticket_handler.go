package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

func (h *TicketHandler) GetTicketsByFlightID(c *gin.Context) {
	isAdmin := c.GetHeader("admin")
	if isAdmin != "true" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Admin privileges required."})
		return
	}

	flightIDStr := c.Query("flightId")
	if flightIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Flight ID is required."})
		return
	}

	flightID, err := strconv.ParseInt(flightIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Flight ID."})
		return
	}

	tickets, err := h.getTicketsByFlightIDUseCase.Execute(c.Request.Context(), flightID)
	if err != nil {
		if err == adapters.ErrFlightNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Flight not found."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error occurred. Please try again later."})
		return
	}

	response := mappers.ToGetTicketsByFlightIDResponse(tickets)
	c.JSON(http.StatusOK, response)
}

func (h *TicketHandler) GetTicket(c *gin.Context) {
	ticketIDStr := c.Query("id")
	if ticketIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id is required"})
		return
	}

	ticketID, err := strconv.ParseInt(ticketIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	ticket, err := h.getTicketUseCase.Execute(c.Request.Context(), ticketID)
	if err != nil {
		if errors.Is(err, adapters.ErrTicketNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Ticket not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error occurred"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Ticket retrieved successfully.",
		"data":    ticket,
	})
}

func (h *TicketHandler) CancelTicket(c *gin.Context) {
	ticketIDStr := c.Query("id")
	if ticketIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id is required"})
		return
	}

	ticketID, err := strconv.ParseInt(ticketIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	ticket, err := h.cancelTicketUseCase.Execute(c.Request.Context(), ticketID)
	if err != nil {
		if errors.Is(err, adapters.ErrTicketNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Ticket not found"})
			return
		}
		if errors.Is(err, adapters.ErrTicketCannotBeCancelled) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Ticket cannot be cancelled due to its current status."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error occurred"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Ticket cancelled successfully.",
		"ticket":  ticket,
	})
}

func (h *TicketHandler) UpdateSeats(c *gin.Context) {
	var updates []dto.UpdateSeatRequest
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid seat data. Please check the input fields."})
		return
	}

	responses, err := h.updateSeatsUseCase.Execute(c.Request.Context(), updates)
	if err != nil {
		if errors.Is(err, adapters.ErrTicketNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "One or more tickets not found."})
			return
		}
		if errors.Is(err, adapters.ErrInvalidSeat) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid seat data. Please check the input fields."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error occurred. Please try again later.", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Seats updated successfully.",
		"data":    mappers.ToUpdateSeatResponses(responses),
	})
}
