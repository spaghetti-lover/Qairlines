package usecases

import (
	"context"
	"time"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

// AirplaneModelCreateUseCase is a use case for creating an airplane model.
type IAirplaneModelCreateUseCase interface {
	Execute(airplaneModelCreateInput AirplaneModelCreateInput) (AirplaneModelCreateOutput, error)
}

// AirplaneModelCreateUseCase implements the IAirplaneModelCreateUseCase interface.
type AirplaneModelCreateUseCase struct {
	airplaneModelRepository adapters.IAirplaneModelRepository
}

// AirplaneModelCreateInput is the input for the AirplaneModelCreateUseCase.
type AirplaneModelCreateInput struct {
	Name         string
	Manufacturer string
	TotalSeats   int64
}

// AirplaneModelCreateOutput is the output for the AirplaneModelCreateUseCase.
type AirplaneModelCreateOutput struct {
	AirplaneModelID int64
	Name            string
	Manufacturer    string
	TotalSeats      int64
	CreatedAt       time.Time
}

// NewAirplaneModelCreateUseCase creates a new instance of AirplaneModelCreateUseCase.
func NewAirplaneModelCreateUseCase(airplaneModelRepository adapters.IAirplaneModelRepository) IAirplaneModelCreateUseCase {
	return &AirplaneModelCreateUseCase{
		airplaneModelRepository: airplaneModelRepository,
	}
}

func (u *AirplaneModelCreateUseCase) Execute(input AirplaneModelCreateInput) (AirplaneModelCreateOutput, error) {
	model, err := entities.NewAirplaneModel(input.Name, input.Manufacturer, input.TotalSeats)
	if err != nil {
		return AirplaneModelCreateOutput{}, err
	}

	airplaneModel, err := u.airplaneModelRepository.CreateAirplaneModel(context.Background(), entities.CreateAirplaneModelParams{
		Name:         model.Name,
		Manufacturer: model.Manufacturer,
		TotalSeats:   model.TotalSeats,
	})
	if err != nil {
		return AirplaneModelCreateOutput{}, err
	}

	return AirplaneModelCreateOutput{
		AirplaneModelID: airplaneModel.AirplaneModelID,
		Name:            airplaneModel.Name,
		Manufacturer:    airplaneModel.Manufacturer,
		TotalSeats:      airplaneModel.TotalSeats,
		CreatedAt:       airplaneModel.CreatedAt,
	}, nil
}
