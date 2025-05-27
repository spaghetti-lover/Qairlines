package ticket

import (
	"context"
	"errors"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type ICancelTicketUseCase interface {
	Execute(ctx context.Context, ticketID int64) (*entities.Ticket, error)
}

type CancelTicketUseCase struct {
	ticketRepository adapters.ITicketRepository
}

func NewCancelTicketUseCase(ticketRepository adapters.ITicketRepository) ICancelTicketUseCase {
	return &CancelTicketUseCase{
		ticketRepository: ticketRepository,
	}
}
func (u *CancelTicketUseCase) Execute(ctx context.Context, ticketID int64) (*entities.Ticket, error) {
	ticket, err := u.ticketRepository.CancelTicket(ctx, ticketID)
	if err != nil {
		if errors.Is(err, adapters.ErrTicketNotFound) {
			return nil, adapters.ErrTicketNotFound
		}
		if errors.Is(err, adapters.ErrTicketCannotBeCancelled) {
			return nil, adapters.ErrTicketCannotBeCancelled
		}
		return nil, err
	}
	return ticket, nil
}
