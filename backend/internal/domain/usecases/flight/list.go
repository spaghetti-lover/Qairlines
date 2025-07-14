package flight

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
)

type IListFlightsUseCase interface {
	Execute(ctx context.Context, page int, limit int) ([]dto.FlightSearchResponse, error)
}

type listFlightsUseCase struct {
	flightRepository adapters.IFlightRepository
}

func NewlistFlightsUseCase(flightRepository adapters.IFlightRepository) IListFlightsUseCase {
	return &listFlightsUseCase{
		flightRepository: flightRepository,
	}
}

func (u *listFlightsUseCase) Execute(ctx context.Context, page int, limit int) ([]dto.FlightSearchResponse, error) {
	start := (page - 1) * limit
	flights, err := u.flightRepository.ListFlights(ctx, start, limit)
	if err != nil {
		if err == adapters.ErrNoSuggestedFlights {
			return nil, nil
		}
		return nil, err
	}

	// Map flights to DTO
	return mappers.ToFlightSearchResponses(flights), nil
}
