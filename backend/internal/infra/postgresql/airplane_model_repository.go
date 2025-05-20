package postgresql

import (
	"context"

	db "github.com/spaghetti-lover/qairlines/db/sqlc"
	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

// AirplaneModelRepositoryPostgres is a struct that implements the IAirplaneModelRepository interface.
type AirplaneModelRepositoryPostgres struct {
	store db.Store
}

// NewAirplaneModelRepositoryPostgres creates a new instance of AirplaneModelRepositoryPostgres.
func NewAirplaneModelRepositoryPostgres(store *db.Store) adapters.IAirplaneModelRepository {
	return &AirplaneModelRepositoryPostgres{store: *store}
}

// NewAirplaneModelRepositoryPostgres creates a new instance of AirplaneModelRepositoryPostgres
func (r *AirplaneModelRepositoryPostgres) CreateAirplaneModel(ctx context.Context, arg entities.CreateAirplaneModelParams) (entities.AirplaneModel, error) {
	airplaneModel, err := r.store.CreateAirplaneModel(ctx, db.CreateAirplaneModelParams{
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
	err := r.store.DeleteAirplaneModel(ctx, airplaneModelID)
	if err != nil {
		return err
	}
	return nil
}

// GetAirplaneModel retrieves an airplane model by its ID.
func (r *AirplaneModelRepositoryPostgres) GetAirplaneModel(ctx context.Context, airplaneModelID int64) (entities.AirplaneModel, error) {
	airplaneModel, err := r.store.GetAirplaneModel(ctx, airplaneModelID)

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
	airplaneModels, err := r.store.ListAirplaneModels(ctx, db.ListAirplaneModelsParams{
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
