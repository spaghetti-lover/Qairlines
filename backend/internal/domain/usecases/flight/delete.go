package flight

import (
	"context"
	"errors"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
)

type IDeleteFlightUseCase interface {
	Execute(ctx context.Context, flightID int64) error
}

type DeleteFlightUseCase struct {
	flightRepository adapters.IFlightRepository
}

func NewDeleteFlightUseCase(flightRepository adapters.IFlightRepository) IDeleteFlightUseCase {
	return &DeleteFlightUseCase{
		flightRepository: flightRepository,
	}
}

func (u *DeleteFlightUseCase) Execute(ctx context.Context, flightID int64) error {
	// Xóa chuyến bay trong repository
	err := u.flightRepository.DeleteFlightByID(ctx, flightID)
	if err != nil {
		if errors.Is(err, adapters.ErrFlightNotFound) {
			return adapters.ErrFlightNotFound
		}
		return err
	}

	return nil
}
