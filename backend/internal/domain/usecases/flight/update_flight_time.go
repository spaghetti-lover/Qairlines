package flight

import (
	"context"
	"errors"
	"time"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
)

type IUpdateFlightTimesUseCase interface {
	Execute(ctx context.Context, flightID int64, request dto.UpdateFlightTimesRequest) (*dto.UpdateFlightTimesResponse, error)
}

type UpdateFlightTimesUseCase struct {
	flightRepository adapters.IFlightRepository
}

func NewUpdateFlightTimesUseCase(flightRepository adapters.IFlightRepository) IUpdateFlightTimesUseCase {
	return &UpdateFlightTimesUseCase{
		flightRepository: flightRepository,
	}
}

func (u *UpdateFlightTimesUseCase) Execute(ctx context.Context, flightID int64, request dto.UpdateFlightTimesRequest) (*dto.UpdateFlightTimesResponse, error) {
	// Chuyển đổi thời gian từ seconds sang time.Time
	departureTime := time.Unix(request.DepartureTime.Seconds, 0)
	arrivalTime := time.Unix(request.ArrivalTime.Seconds, 0)

	// Cập nhật thời gian chuyến bay trong repository
	flight, err := u.flightRepository.UpdateFlightTimes(ctx, flightID, departureTime, arrivalTime)
	if err != nil {
		if errors.Is(err, adapters.ErrFlightNotFound) {
			return nil, adapters.ErrFlightNotFound
		}
		return nil, err
	}

	// Sử dụng mapper để chuyển đổi entity sang DTO
	return mappers.ToUpdateFlightTimesResponse(flight), nil
}
