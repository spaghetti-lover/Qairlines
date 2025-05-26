package ticket

import (
	"context"
	"errors"
	"strconv"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
)

var ErrTicketNotFound = errors.New("ticket not found")

type IGetTicketUseCase interface {
	Execute(ctx context.Context, ticketID int64) (*dto.GetTicketResponse, error)
}

type GetTicketUseCase struct {
	ticketRepository adapters.ITicketRepository
}

func NewGetTicketUseCase(ticketRepository adapters.ITicketRepository) IGetTicketUseCase {
	return &GetTicketUseCase{
		ticketRepository: ticketRepository,
	}
}

func (u *GetTicketUseCase) Execute(ctx context.Context, ticketID int64) (*dto.GetTicketResponse, error) {
	// Lấy thông tin vé từ repository
	ticket, err := u.ticketRepository.GetTicketByID(ctx, ticketID)
	if err != nil {
		if errors.Is(err, adapters.ErrTicketNotFound) {
			return nil, ErrTicketNotFound
		}
		return nil, err
	}

	// Map entity sang DTO
	return &dto.GetTicketResponse{
		TicketID:    ticket.TicketID,
		Status:      string(ticket.Status),
		SeatCode:    ticket.SeatCode,
		FlightClass: string(ticket.FlightClass),
		Price:       ticket.Price,
		OwnerData: struct {
			FirstName   string `json:"firstName"`
			LastName    string `json:"lastName"`
			Gender      string `json:"gender"`
			PhoneNumber string `json:"phoneNumber"`
		}{
			FirstName:   ticket.Owner.FirstName,
			LastName:    ticket.Owner.LastName,
			Gender:      ticket.Owner.Gender,
			PhoneNumber: ticket.Owner.PhoneNumber,
		},
		BookingID: strconv.FormatInt(ticket.BookingID, 10),
		FlightID:  strconv.FormatInt(ticket.FlightID, 10),
		CreatedAt: ticket.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt: ticket.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}, nil
}
