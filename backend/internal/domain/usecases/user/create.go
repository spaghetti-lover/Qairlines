package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

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

func (u *UserCreateUseCase) Execute(ctx context.Context, arg entities.CreateUserParams) (entities.User, error) {
	// Kiểm tra email đã tồn tại
	existingUser, err := u.userRepository.GetUserByEmail(ctx, arg.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return entities.User{}, fmt.Errorf("internal error: %w", err)
	}
	if existingUser != nil {
		return entities.User{}, fmt.Errorf("email already in use")
	}

	user, err := u.userRepository.CreateUser(ctx, arg)
	if err != nil {
		return entities.User{}, fmt.Errorf("failed to create user: %w", err)
	}
	return user, nil
}
