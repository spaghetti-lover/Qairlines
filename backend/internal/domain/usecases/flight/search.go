package flight

import (
	"context"
	"errors"
	"time"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
)

type ISearchFlightsUseCase interface {
	Execute(ctx context.Context, departureCity, arrivalCity string, flightDate string) ([]dto.FlightSearchResponse, error)
}

type SearchFlightsUseCase struct {
	flightRepository adapters.IFlightRepository
}

func NewSearchFlightsUseCase(flightRepository adapters.IFlightRepository) ISearchFlightsUseCase {
	return &SearchFlightsUseCase{
		flightRepository: flightRepository,
	}
}

func (u *SearchFlightsUseCase) Execute(ctx context.Context, departureCity, arrivalCity string, flightDate string) ([]dto.FlightSearchResponse, error) {
	// Parse flightDate to time.Time
	date, err := time.Parse("2006-01-02", flightDate)
	if err != nil {
		return nil, errors.New("invalid query parameters. Please check departureCity, arrivalCity, and flightDate")
	}

	// Search flights in repository
	flights, err := u.flightRepository.SearchFlights(ctx, departureCity, arrivalCity, date)
	if err != nil {
		if errors.Is(err, adapters.ErrNoFlightsFound) {
			return nil, adapters.ErrNoFlightsFound
		}
		return nil, err
	}

	// Map flights to DTO
	return mappers.ToFlightSearchResponses(flights), nil
}
