package admin

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
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
	if admin.UserID == 0 {
		return entities.Admin{}, adapters.ErrAdminNotFound
	}
	return admin, nil
}
