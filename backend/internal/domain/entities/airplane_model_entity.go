package entities

import (
	"errors"
	"time"
)

// AirplaneModel represents an airplane model with its details.
type AirplaneModel struct {
	AirplaneModelID int64
	Name            string
	Manufacturer    string
	TotalSeats      int64
	CreatedAt       time.Time
}

// CreateAirplaneModelParams contains the parameters required to create a new AirplaneModel.
type CreateAirplaneModelParams struct {
	Name         string
	Manufacturer string
	TotalSeats   int64
}

// AirplaneModelRepository represents the interface for airplane model repository operations.
type ListAirplaneModelsParams struct {
	Limit  int32
	Offset int32
}

// NewAirplaneModel creates a new AirplaneModel instance with the given parameters.
func NewAirplaneModel(name, manufacturer string, totalSeats int64) (*AirplaneModel, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}
	if manufacturer == "" {
		return nil, errors.New("manufacturer is required")
	}
	if totalSeats <= 0 {
		return nil, errors.New("totalSeats must be positive")
	}
	return &AirplaneModel{
		Name:         name,
		Manufacturer: manufacturer,
		TotalSeats:   totalSeats,
		CreatedAt:    time.Now(),
	}, nil
}
