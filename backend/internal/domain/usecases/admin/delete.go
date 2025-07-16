package admin

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
)

type IDeleteAdminUseCase interface {
	Execute(ctx context.Context, userID int64) error
}
type DeleteAdminUseCase struct {
	adminRepository adapters.IAdminRepository
}

func NewDeleteAdminUseCase(adminRepository adapters.IAdminRepository) IDeleteAdminUseCase {
	return &DeleteAdminUseCase{
		adminRepository: adminRepository,
	}
}

func (u *DeleteAdminUseCase) Execute(ctx context.Context, userID int64) error {
	admin, err := u.adminRepository.GetAdminByUserID(ctx, userID)
	if err != nil {
		return err
	}

	if admin.UserID == 0 {
		return adapters.ErrAdminNotFound
	}

	return u.adminRepository.DeleteAdmin(ctx, userID)
}
