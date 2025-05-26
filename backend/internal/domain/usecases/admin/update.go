package admin

import (
	"context"
	"errors"
	"strconv"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type UpdateAdminInput struct {
	UserID    int64
	FirstName string
	LastName  string
	Email     string
}

type IUpdateAdminUseCase interface {
	Execute(ctx context.Context, input UpdateAdminInput) (entities.Admin, error)
}

type UpdateAdminUseCase struct {
	adminRepository adapters.IAdminRepository
	userRepository  adapters.IUserRepository
}

func NewUpdateAdminUseCase(adminRepository adapters.IAdminRepository, userRepository adapters.IUserRepository) IUpdateAdminUseCase {
	return &UpdateAdminUseCase{
		adminRepository: adminRepository,
		userRepository:  userRepository,
	}
}

func (u *UpdateAdminUseCase) Execute(ctx context.Context, input UpdateAdminInput) (entities.Admin, error) {
	adminID := strconv.FormatInt(input.UserID, 10)
	if adminID == "" {
		return entities.Admin{}, errors.New("invalid admin ID")
	}
	admin, err := u.adminRepository.GetAdminByUserID(ctx, input.UserID)
	if err != nil {
		return entities.Admin{}, ErrAdminNotFound
	}

	admin.FirstName = input.FirstName
	admin.LastName = input.LastName
	admin.Email = input.Email

	updatedAdmin, err := u.adminRepository.UpdateAdmin(ctx, admin)
	if err != nil {
		return entities.Admin{}, err
	}

	return updatedAdmin, nil
}
