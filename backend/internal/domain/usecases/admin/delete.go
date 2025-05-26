package admin

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
)

// Interface cho use case
type IDeleteAdminUseCase interface {
	Execute(ctx context.Context, userID int64) error
}

// Implement use case
type DeleteAdminUseCase struct {
	adminRepository adapters.IAdminRepository
}

// Constructor
func NewDeleteAdminUseCase(adminRepository adapters.IAdminRepository) IDeleteAdminUseCase {
	return &DeleteAdminUseCase{
		adminRepository: adminRepository,
	}
}

// Execute thực hiện việc xóa admin
func (u *DeleteAdminUseCase) Execute(ctx context.Context, userID int64) error {
	admin, err := u.adminRepository.GetAdminByUserID(ctx, userID)
	if err != nil {
		return err
	}

	if admin.UserID == "" {
		return ErrAdminNotFound
	}

	// Thực hiện xóa admin
	return u.adminRepository.DeleteAdmin(ctx, userID)
}
