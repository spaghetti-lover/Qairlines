package mappers

import (
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
)

func TicketEntityToResponse(ticket entities.Ticket) dto.TicketResponse {
	return dto.TicketResponse{
		TicketID:    ticket.TicketID,
		Status:      string(ticket.Status),
		SeatCode:    ticket.Seat.SeatCode,
		FlightClass: ticket.Seat.Class,
		Price:       ticket.Price,
		OwnerData: dto.TicketOwnerData{
			FirstName:   ticket.Owner.FirstName,
			LastName:    ticket.Owner.LastName,
			PhoneNumber: ticket.Owner.PhoneNumber,
			Gender:      ticket.Owner.Gender,
		},
		BookingID: ticket.BookingID,
		FlightID:  ticket.FlightID,
		CreatedAt: ticket.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt: ticket.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}
}

func TicketsEntitiesToResponse(tickets []entities.Ticket) []dto.TicketResponse {
	responses := make([]dto.TicketResponse, len(tickets))
	for i, ticket := range tickets {
		responses[i] = TicketEntityToResponse(ticket)
	}
	return responses
}
