package usecases

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type IHealthUseCase interface {
	Execute() (entities.Health, error)
}

// HealthUseCase implements the IHealthUseCase interface.
type HealthUseCase struct {
	healthRepository adapters.IHealthRepository
}

// NewHealthUseCase creates a new instance of HealthUseCase.
func NewHealthUseCase(healthRepository adapters.IHealthRepository) IHealthUseCase {
	return &HealthUseCase{
		healthRepository: healthRepository,
	}
}

func (u *HealthUseCase) Execute() (entities.Health, error) {
	health, err := u.healthRepository.GetHealth(context.Background())
	if err != nil {
		return entities.Health{}, err
	}
	return health, nil
}
