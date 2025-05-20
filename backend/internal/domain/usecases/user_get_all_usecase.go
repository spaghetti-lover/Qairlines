package usecases

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type UserGetOutput struct {
	UserID   int64
	Name     string
	Password string
	Role     entities.UserRole
}

type IUserGetAllUseCase interface {
	Execute(ctx context.Context) ([]entities.User, error)
}

type UserGetAllUseCase struct {
	userRepository adapters.IUserRepository
}

func NewUserGetAllUseCase(userRepository adapters.IUserRepository) IUserGetAllUseCase {
	return &UserGetAllUseCase{
		userRepository: userRepository,
	}
}

func (r *UserGetAllUseCase) Execute(ctx context.Context) ([]entities.User, error) {
	users, err := r.userRepository.GetAllUser(ctx)
	if err != nil {
		return []entities.User{}, err
	}
	return users, nil
}
