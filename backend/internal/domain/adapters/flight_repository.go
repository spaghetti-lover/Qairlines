package adapters

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type IFlightRepository interface {
	CreateFlight(ctx context.Context, flight entities.Flight) (entities.Flight, error)
}
