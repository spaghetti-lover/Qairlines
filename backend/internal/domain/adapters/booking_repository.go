package adapters

import (
	"context"
	"errors"

	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

var (
	ErrInvalidBooking = errors.New("invalid booking data")
)

type IBookingRepository interface {
	CreateBookingTx(ctx context.Context, booking entities.CreateBookingParams) (entities.Booking, []entities.Ticket, []entities.Ticket, error)
}
