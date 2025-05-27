package flight

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
)

type IGetSuggestedFlightsUseCase interface {
	Execute(ctx context.Context) ([]dto.FlightSearchResponse, error)
}

type GetSuggestedFlightsUseCase struct {
	flightRepository adapters.IFlightRepository
}

func NewGetSuggestedFlightsUseCase(flightRepository adapters.IFlightRepository) IGetSuggestedFlightsUseCase {
	return &GetSuggestedFlightsUseCase{
		flightRepository: flightRepository,
	}
}

func (u *GetSuggestedFlightsUseCase) Execute(ctx context.Context) ([]dto.FlightSearchResponse, error) {
	flights, err := u.flightRepository.GetSuggestedFlights(ctx)
	if err != nil {
		if err == adapters.ErrNoSuggestedFlights {
			return nil, nil
	}
		return nil, err
	}

	// Map flights to DTO
	return mappers.ToFlightSearchResponses(flights), nil
}
