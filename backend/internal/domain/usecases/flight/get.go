package flight

import (
	"context"
	"errors"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
)

type IGetFlightUseCase interface {
	Execute(ctx context.Context, flightID int64) (*dto.GetFlightResponse, error)
}

type GetFlightUseCase struct {
	flightRepository adapters.IFlightRepository
}

func NewGetFlightUseCase(flightRepository adapters.IFlightRepository) IGetFlightUseCase {
	return &GetFlightUseCase{
		flightRepository: flightRepository,
	}
}

func (u *GetFlightUseCase) Execute(ctx context.Context, flightID int64) (*dto.GetFlightResponse, error) {
	// Lấy thông tin chuyến bay từ repository
	flight, err := u.flightRepository.GetFlightByID(ctx, flightID)
	if err != nil {
		if errors.Is(err, adapters.ErrFlightNotFound) {
			return nil, adapters.ErrFlightNotFound
		}
		return nil, err
	}

	// Map entity sang DTO
	return mappers.MapFlightToGetFlightResponse(flight), nil
}
