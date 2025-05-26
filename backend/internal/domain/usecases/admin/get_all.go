package admin

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type IGetAllAdminsUseCase interface {
	Execute(ctx context.Context) ([]entities.Admin, error)
}

type GetAllAdminsUseCase struct {
	adminRepository adapters.IAdminRepository
}

func NewGetAllAdminsUseCase(adminRepository adapters.IAdminRepository) IGetAllAdminsUseCase {
	return &GetAllAdminsUseCase{
		adminRepository: adminRepository,
	}
}

func (u *GetAllAdminsUseCase) Execute(ctx context.Context) ([]entities.Admin, error) {
	return u.adminRepository.GetAllAdmins(ctx)
}
