package admin

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type IListAdminsUseCase interface {
	Execute(ctx context.Context, page int, limit int) ([]entities.Admin, error)
}

type ListAdminsUseCase struct {
	adminRepository adapters.IAdminRepository
}

func NewListAdminsUseCase(adminRepository adapters.IAdminRepository) IListAdminsUseCase {
	return &ListAdminsUseCase{
		adminRepository: adminRepository,
	}
}

func (u *ListAdminsUseCase) Execute(ctx context.Context, page int, limit int) ([]entities.Admin, error) {
	start := (page - 1) * limit
	return u.adminRepository.ListAdmins(ctx, start, limit)
}
