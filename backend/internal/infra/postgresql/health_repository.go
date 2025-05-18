package postgresql

import (
	"context"
	"runtime"
	"time"

	db "github.com/spaghetti-lover/qairlines/db/sqlc"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
	"github.com/spaghetti-lover/qairlines/pkg/utils"
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
	newCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	monitor := &utils.CPUMonitor{}
	cpuPercent, err := monitor.Check(newCtx)
	if err != nil {
		return entities.Health{}, err
	}

	return entities.Health{
		Status:  "OK",
		Version: "1.0.0",
		Stats: entities.Stats{
			CPUPercent: cpuPercent,
			CPUCore:    runtime.NumCPU(),
		},
	}, nil
}
