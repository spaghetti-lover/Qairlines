package ticket

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type IGetTicketsByFlightIDUseCase interface {
	Execute(ctx context.Context, flightID int64) ([]entities.Ticket, error)
}

type GetTicketsByFlightIDUseCase struct {
	ticketRepository adapters.ITicketRepository
}

func NewGetTicketsByFlightIDUseCase(ticketRepository adapters.ITicketRepository) IGetTicketsByFlightIDUseCase {
	return &GetTicketsByFlightIDUseCase{
		ticketRepository: ticketRepository,
	}
}

func (u *GetTicketsByFlightIDUseCase) Execute(ctx context.Context, flightID int64) ([]entities.Ticket, error) {
	tickets, err := u.ticketRepository.GetTicketsByFlightID(ctx, flightID)
	if err != nil {
		return nil, err
	}
	if len(tickets) == 0 {
		return nil, adapters.ErrFlightNotFound
	}
	return tickets, err
}
