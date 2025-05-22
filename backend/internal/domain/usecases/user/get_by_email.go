package user

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type IUserGetByEmailUseCase interface {
	Execute(ctx context.Context, email string) (*entities.User, error)
}
type UserGetByEmailUseCase struct {
	userRepository adapters.IUserRepository
}

func NewUserGetByEmailUseCase(userRepository adapters.IUserRepository) IUserGetByEmailUseCase {
	return &UserGetByEmailUseCase{
		userRepository: userRepository,
	}
}
func (r *UserGetByEmailUseCase) Execute(ctx context.Context, email string) (*entities.User, error) {
	user, err := r.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
