package flight

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type IGetAllFlightsUseCase interface {
	Execute(ctx context.Context) ([]entities.Flight, error)
}

type GetAllFlightsUseCase struct {
	flightRepository adapters.IFlightRepository
}

func NewGetAllFlightsUseCase(flightRepository adapters.IFlightRepository) IGetAllFlightsUseCase {
	return &GetAllFlightsUseCase{
		flightRepository: flightRepository,
	}
}

func (u *GetAllFlightsUseCase) Execute(ctx context.Context) ([]entities.Flight, error) {
	// Lấy danh sách chuyến bay từ repository
	flights, err := u.flightRepository.GetAllFlights(ctx)
	if err != nil {
		return nil, err
	}

	// Sử dụng mapper để chuyển đổi danh sách entity sang DTO
	return flights, nil
}
