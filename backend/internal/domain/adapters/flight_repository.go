package adapters

import (
	"context"
	"errors"
	"time"

	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

var (
	ErrFlightNotFound     = errors.New("flight not found")
	ErrNoFlightsFound     = errors.New("no flights found for the given criteria")
	ErrNoSuggestedFlights = errors.New("no suggested flights available")
)

type IFlightRepository interface {
	CreateFlight(ctx context.Context, flight entities.Flight) (entities.Flight, error)
	GetFlightByID(ctx context.Context, flightID int64) (*entities.Flight, error)
	UpdateFlightTimes(ctx context.Context, flightID int64, departureTime, arrivalTime time.Time) (*entities.Flight, error)
	GetAllFlights(ctx context.Context) ([]entities.Flight, error)
	DeleteFlightByID(ctx context.Context, flightID int64) error
	SearchFlights(ctx context.Context, departureCity, arrivalCity string, flightDate time.Time) ([]entities.Flight, error)
	GetSuggestedFlights(ctx context.Context) ([]entities.Flight, error)
}
