package ticket

import (
	"context"
	"errors"
	"fmt"

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
	// 1. Kiểm tra xem vé có tồn tại không
	ticket, err := u.ticketRepository.GetTicketByID(ctx, ticketID)
	if err != nil {
		if errors.Is(err, adapters.ErrTicketNotFound) {
			return nil, adapters.ErrTicketNotFound
		}
		return nil, fmt.Errorf("error fetching ticket: %w", err)
	}

	// 2. Kiểm tra xem vé có thể hủy được không
	if ticket.Status != entities.TicketStatusBooked {
		return nil, adapters.ErrTicketCannotBeCancelled
	}

	// 3. Hủy vé
	updatedTicket, err := u.ticketRepository.CancelTicket(ctx, ticketID)
	if err != nil {
		return nil, fmt.Errorf("error cancelling ticket: %w", err)
	}

	return updatedTicket, nil
}
