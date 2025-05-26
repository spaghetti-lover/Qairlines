package flight

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type ICreateFlightUseCase interface {
	Execute(ctx context.Context, flight entities.Flight) (entities.Flight, error)
}

type CreateFlightUseCase struct {
	flightRepository adapters.IFlightRepository
}

func NewCreateFlightUseCase(flightRepository adapters.IFlightRepository) ICreateFlightUseCase {
	return &CreateFlightUseCase{flightRepository: flightRepository}
}

func (u *CreateFlightUseCase) Execute(ctx context.Context, flight entities.Flight) (entities.Flight, error) {
	return u.flightRepository.CreateFlight(ctx, flight)
}
