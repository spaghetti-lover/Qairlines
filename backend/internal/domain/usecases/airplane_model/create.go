package airplane_model

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

// AirplaneModelCreateUseCase is a use case for creating an airplane model.
type IAirplaneModelCreateUseCase interface {
	Execute(ctx context.Context, airplaneModelCreateInput entities.CreateAirplaneModelParams) (entities.AirplaneModel, error)
}

// AirplaneModelCreateUseCase implements the IAirplaneModelCreateUseCase interface.
type AirplaneModelCreateUseCase struct {
	airplaneModelRepository adapters.IAirplaneModelRepository
}

// NewAirplaneModelCreateUseCase creates a new instance of AirplaneModelCreateUseCase.
func NewAirplaneModelCreateUseCase(airplaneModelRepository adapters.IAirplaneModelRepository) IAirplaneModelCreateUseCase {
	return &AirplaneModelCreateUseCase{
		airplaneModelRepository: airplaneModelRepository,
	}
}

func (u *AirplaneModelCreateUseCase) Execute(ctx context.Context, arg entities.CreateAirplaneModelParams) (entities.AirplaneModel, error) {
	model, err := entities.NewAirplaneModel(arg.Name, arg.Manufacturer, arg.TotalSeats)
	if err != nil {
		return entities.AirplaneModel{}, err
	}

	airplaneModel, err := u.airplaneModelRepository.CreateAirplaneModel(context.Background(), entities.CreateAirplaneModelParams{
		Name:         model.Name,
		Manufacturer: model.Manufacturer,
		TotalSeats:   model.TotalSeats,
	})
	if err != nil {
		return entities.AirplaneModel{}, err
	}

	return airplaneModel, nil
}
