package postgresql

import (
	"context"

	db "github.com/spaghetti-lover/qairlines/db/sqlc"
	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

// AirplaneModelRepositoryPostgres is a struct that implements the IAirplaneModelRepository interface.
type AirplaneModelRepositoryPostgres struct {
	queries *db.Queries
}

// NewAirplaneModelRepositoryPostgres creates a new instance of AirplaneModelRepositoryPostgres.
func NewAirplaneModelRepositoryPostgres() adapters.IAirplaneModelRepository {
	return &AirplaneModelRepositoryPostgres{}
}

// NewAirplaneModelRepositoryPostgres creates a new instance of AirplaneModelRepositoryPostgres
func (r *AirplaneModelRepositoryPostgres) CreateAirplaneModel(ctx context.Context, arg entities.CreateAirplaneModelParams) (entities.AirplaneModel, error) {
	airplaneModel, err := r.queries.CreateAirplaneModel(ctx, db.CreateAirplaneModelParams{
		Name:         arg.Name,
		Manufacturer: arg.Manufacturer,
		TotalSeats:   arg.TotalSeats,
	})

	if err != nil {
		return entities.AirplaneModel{}, err
	}

	return entities.AirplaneModel{
		AirplaneModelID: airplaneModel.AirplaneModelID,
		Name:            airplaneModel.Name,
		Manufacturer:    airplaneModel.Manufacturer,
		TotalSeats:      airplaneModel.TotalSeats,
		CreatedAt:       airplaneModel.CreatedAt,
	}, nil
}

// DeleteAirplaneModel deletes an airplane model by its ID.
func (r *AirplaneModelRepositoryPostgres) DeleteAirplaneModel(ctx context.Context, airplaneModelID int64) error {
	err := r.queries.DeleteAirplaneModel(ctx, airplaneModelID)
	if err != nil {
		return err
	}
	return nil
}

// GetAirplaneModel retrieves an airplane model by its ID.
func (r *AirplaneModelRepositoryPostgres) GetAirplaneModel(ctx context.Context, airplaneModelID int64) (entities.AirplaneModel, error) {
	airplaneModel, err := r.queries.GetAirplaneModel(ctx, airplaneModelID)

	if err != nil {
		return entities.AirplaneModel{}, err
	}

	return entities.AirplaneModel{
		AirplaneModelID: airplaneModel.AirplaneModelID,
		Name:            airplaneModel.Name,
		Manufacturer:    airplaneModel.Manufacturer,
		TotalSeats:      airplaneModel.TotalSeats,
		CreatedAt:       airplaneModel.CreatedAt,
	}, nil
}

// ListAirplaneModels retrieves a list of airplane models with pagination.
func (r *AirplaneModelRepositoryPostgres) ListAirplaneModels(ctx context.Context, arg entities.ListAirplaneModelsParams) ([]entities.AirplaneModel, error) {
	airplaneModels, err := r.queries.ListAirplaneModels(ctx, db.ListAirplaneModelsParams{
		Limit:  arg.Limit,
		Offset: arg.Offset,
	})

	if err != nil {
		return []entities.AirplaneModel{}, err
	}

	result := []entities.AirplaneModel{}
	for _, airplaneModel := range airplaneModels {
		result = append(result, entities.AirplaneModel{
			AirplaneModelID: airplaneModel.AirplaneModelID,
			Name:            airplaneModel.Name,
			Manufacturer:    airplaneModel.Manufacturer,
			TotalSeats:      airplaneModel.TotalSeats,
			CreatedAt:       airplaneModel.CreatedAt,
		})
	}
	return result, nil
}
