package adapters

import (
	"context"
	"errors"
	"time"

	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

var ErrFlightNotFound = errors.New("flight not found")

type IFlightRepository interface {
	CreateFlight(ctx context.Context, flight entities.Flight) (entities.Flight, error)
	GetFlightByID(ctx context.Context, flightID int64) (*entities.Flight, error)
	UpdateFlightTimes(ctx context.Context, flightID int64, departureTime, arrivalTime time.Time) (*entities.Flight, error)
}
