package adapters

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type IHealthRepository interface {
	GetHealth(ctx context.Context) (entities.Health, error)
}
