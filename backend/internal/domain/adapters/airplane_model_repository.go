package adapters

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

// IAirplaneModelRepository is an interface that defines the methods for interacting with airplane models in the repository.
type IAirplaneModelRepository interface {
	CreateAirplaneModel(ctx context.Context, arg entities.CreateAirplaneModelParams) (entities.AirplaneModel, error)
	DeleteAirplaneModel(ctx context.Context, airplaneModelID int64) error
	GetAirplaneModel(ctx context.Context, airplaneModelID int64) (entities.AirplaneModel, error)
	ListAirplaneModels(ctx context.Context, arg entities.ListAirplaneModelsParams) ([]entities.AirplaneModel, error)
}
