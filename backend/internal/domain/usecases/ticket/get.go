package ticket

import (
	"context"
	"errors"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type IGetTicketUseCase interface {
	Execute(ctx context.Context, ticketID int64) (*entities.Ticket, error)
}

type GetTicketUseCase struct {
	ticketRepository adapters.ITicketRepository
}

func NewGetTicketUseCase(ticketRepository adapters.ITicketRepository) IGetTicketUseCase {
	return &GetTicketUseCase{
		ticketRepository: ticketRepository,
	}
}

func (u *GetTicketUseCase) Execute(ctx context.Context, ticketID int64) (*entities.Ticket, error) {

	ticket, err := u.ticketRepository.GetTicketByID(ctx, ticketID)
	if err != nil {
		if errors.Is(err, adapters.ErrTicketNotFound) {
			return nil, adapters.ErrTicketNotFound
		}
		return nil, err
	}

	return ticket, nil
}
