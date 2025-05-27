package mappers

import (
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
)

func ToGetTicketsByFlightIDResponse(tickets []entities.Ticket) []dto.GetTicketResponse {
	var responses []dto.GetTicketResponse

	for _, ticket := range tickets {
		responses = append(responses, dto.GetTicketResponse{
			TicketID:    ticket.TicketID,
			SeatID:      ticket.SeatID,
			FlightClass: string(ticket.FlightClass),
			Price:       ticket.Price,
			Status:      string(ticket.Status),
			BookingID:   ticket.BookingID,
			FlightID:    ticket.FlightID,
			CreatedAt:   ticket.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt:   ticket.UpdatedAt.Format("2006-01-02T15:04:05Z"),
			Seat: dto.SeatResponse{
				SeatID:      ticket.Seat.SeatID,
				SeatCode:    ticket.Seat.SeatCode,
				IsAvailable: ticket.Seat.IsAvailable,
				Class:       string(ticket.Seat.Class),
			},
			Owner: dto.TicketOwnerResponse{
				FirstName:   ticket.Owner.FirstName,
				LastName:    ticket.Owner.LastName,
				PhoneNumber: ticket.Owner.PhoneNumber,
				Gender:      string(ticket.Owner.Gender),
			},
		})
	}

	return responses
}
