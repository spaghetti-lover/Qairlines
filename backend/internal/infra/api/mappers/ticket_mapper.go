package mappers

import (
	"strconv"

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

func TicketToTicketDetails(ticket *entities.Ticket) dto.TicketDetails {
	return dto.TicketDetails{
		TicketID:    ticket.TicketID,
		Status:      string(ticket.Status),
		SeatCode:    ticket.SeatCode,
		FlightClass: string(ticket.FlightClass),
		Price:       ticket.Price,
		OwnerData: dto.OwnerData{
			FirstName:   ticket.Owner.FirstName,
			LastName:    ticket.Owner.LastName,
			PhoneNumber: ticket.Owner.PhoneNumber,
		},
		BookingID: strconv.FormatInt(ticket.BookingID, 10),
		FlightID:  strconv.FormatInt(ticket.FlightID, 10),
		UpdatedAt: ticket.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}
}

func ToCancelTicketResponse(ticket *entities.Ticket) dto.CancelTicketResponse {
	return dto.CancelTicketResponse{
		Message: "Ticket cancelled successfully.",
		Ticket:  TicketToTicketDetails(ticket),
	}
}
