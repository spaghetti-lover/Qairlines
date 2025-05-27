package mappers

import (
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
)

func ToGetTicketsByFlightIDResponse(tickets []entities.Ticket) []dto.GetTicketByFlightIDResponse {
	var responses []dto.GetTicketByFlightIDResponse

	for _, ticket := range tickets {
		responses = append(responses, dto.GetTicketByFlightIDResponse{
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

func ToGetTicketResponse(ticket entities.Ticket) dto.GetTicketResponse {
	return dto.GetTicketResponse{
		TicketID:    ticket.TicketID,
		Status:      string(ticket.Status),
		SeatCode:    ticket.Seat.SeatCode,
		FlightClass: string(ticket.FlightClass),
		Price:       ticket.Price,
		OwnerData: dto.TicketOwnerResponse{
			FirstName:   ticket.Owner.FirstName,
			LastName:    ticket.Owner.LastName,
			PhoneNumber: ticket.Owner.PhoneNumber,
			Gender:      string(ticket.Owner.Gender),
		},
		BookingID: ticket.BookingID,
		FlightID:  ticket.FlightID,
		CreatedAt: ticket.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt: ticket.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}
}

func ToCancelTicketResponse(ticket *entities.Ticket) *dto.CancelTicketResponse {
	return &dto.CancelTicketResponse{
		TicketID:    ticket.TicketID,
		Status:      string(ticket.Status),
		SeatCode:    ticket.Seat.SeatCode,
		FlightClass: string(ticket.FlightClass),
		Price:       ticket.Price,
		OwnerData: dto.TicketOwnerResponse{
			FirstName:   ticket.Owner.FirstName,
			LastName:    ticket.Owner.LastName,
			PhoneNumber: ticket.Owner.PhoneNumber,
		},
		BookingID: ticket.BookingID,
		FlightID:  ticket.FlightID,
		UpdatedAt: ticket.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}
}

func ToUpdateSeatResponses(tickets []entities.Ticket) []dto.UpdateSeatResponse {
	var responses []dto.UpdateSeatResponse

	for _, ticket := range tickets {
		responses = append(responses, dto.UpdateSeatResponse{
			TicketID: ticket.TicketID,
			SeatCode: ticket.Seat.SeatCode,
			Status:   "Updated",
		})
	}

	return responses
}
