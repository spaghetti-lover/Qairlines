package postgresql

import (
	"context"

	db "github.com/spaghetti-lover/qairlines/db/sqlc"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type HealthRepositoryPostgres struct {
	store *db.Store
}

func NewHealthRepositoryPostgres(store *db.Store) *HealthRepositoryPostgres {
	return &HealthRepositoryPostgres{
		store: store,
	}
}

func (r *HealthRepositoryPostgres) GetHealth(ctx context.Context) (entities.Health, error) {
	return entities.Health{Status: "OK"}, nil
}
