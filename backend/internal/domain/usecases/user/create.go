package user

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type IUserCreateUseCase interface {
	Execute(ctx context.Context, arg entities.CreateUserParams) (entities.User, error)
}

type UserCreateUseCase struct {
	userRepository adapters.IUserRepository
}

func NewUserCreateUseCase(userRepository adapters.IUserRepository) IUserCreateUseCase {
	return &UserCreateUseCase{
		userRepository: userRepository,
	}
}

func (r *UserCreateUseCase) Execute(ctx context.Context, arg entities.CreateUserParams) (entities.User, error) {
	user, err := r.userRepository.CreateUser(ctx, arg)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}
