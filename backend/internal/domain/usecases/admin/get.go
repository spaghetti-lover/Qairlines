package admin

import (
	"context"
	"errors"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

var (
	ErrAdminNotFound = errors.New("admin not found")
)

type IGetCurrentAdminUseCase interface {
	Execute(ctx context.Context, userID int64) (entities.Admin, error)
}

type GetCurrentAdminUseCase struct {
	adminRepository adapters.IAdminRepository
}

func NewGetCurrentAdminUseCase(adminRepository adapters.IAdminRepository) IGetCurrentAdminUseCase {
	return &GetCurrentAdminUseCase{
		adminRepository: adminRepository,
	}
}

func (u *GetCurrentAdminUseCase) Execute(ctx context.Context, userID int64) (entities.Admin, error) {
	admin, err := u.adminRepository.GetAdminByUserID(ctx, userID)
	if err != nil {
		return entities.Admin{}, err
	}
	if admin.UserID == "" {
		return entities.Admin{}, ErrAdminNotFound
	}
	return admin, nil
}
