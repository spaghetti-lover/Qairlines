package adapters

import (
	"context"
	"errors"

	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

var ErrBookingNotFound = errors.New("booking not found")

type IBookingRepository interface {
	GetBookingByID(ctx context.Context, bookingID int64) (*entities.Booking, error)
}
