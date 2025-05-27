package adapters

import (
	"context"
	"errors"

	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

var (
	ErrTicketNotFound          = errors.New("ticket not found")
	ErrTicketCannotBeCancelled = errors.New("ticket cannot be cancelled due to its current status")
	ErrInvalidSeat             = errors.New("invalid seat data")
)

type ITicketRepository interface {
	GetTicketsByFlightID(ctx context.Context, flightID int64) ([]entities.Ticket, error)
	GetTicketByID(ctx context.Context, ticketID int64) (*entities.Ticket, error)
	CancelTicket(ctx context.Context, ticketID int64) (*entities.Ticket, error)
	UpdateSeat(ctx context.Context, ticketID int64, seatCode string) (*entities.Ticket, error)
}
