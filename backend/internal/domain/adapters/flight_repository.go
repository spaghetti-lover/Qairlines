package adapters

import (
	"context"
	"errors"

	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

var ErrFlightNotFound = errors.New("flight not found")

type IFlightRepository interface {
	CreateFlight(ctx context.Context, flight entities.Flight) (entities.Flight, error)
	GetFlightByID(ctx context.Context, flightID int64) (*entities.Flight, error)
}
