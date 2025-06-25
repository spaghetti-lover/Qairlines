package admin

import (
	"context"
	"database/sql"
	"errors"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type ICreateAdminUseCase interface {
	Execute(ctx context.Context, admin entities.CreateUserParams) (entities.User, error)
}

type CreateAdminUseCase struct {
	adminRepository adapters.IAdminRepository
	userRepository  adapters.IUserRepository
}

func NewCreateAdminUseCase(adminRepository adapters.IAdminRepository, userRepository adapters.IUserRepository) ICreateAdminUseCase {
	return &CreateAdminUseCase{
		adminRepository: adminRepository,
		userRepository:  userRepository,
	}
}

func (c *CreateAdminUseCase) Execute(ctx context.Context, admin entities.CreateUserParams) (entities.User, error) {
	// Check if the user already exists
	existingUser, err := c.userRepository.GetUserByEmail(ctx, admin.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return entities.User{}, errors.New("failed to check existing user")
	}
	if existingUser != nil {
		return entities.User{}, errors.New("email already in use")
	}
	createdAdmin, err := c.adminRepository.CreateAdminTx(ctx, admin)
	if err != nil {
		return entities.User{}, err
	}
	return createdAdmin, nil
}
