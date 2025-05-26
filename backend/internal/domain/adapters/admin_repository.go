package adapters

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type IAdminRepository interface {
	CreateAdminTx(ctx context.Context, arg entities.CreateUserParams) (entities.User, error)
	GetAllAdmins(ctx context.Context) ([]entities.Admin, error)
	GetAdminByID(ctx context.Context, adminID string) (entities.Admin, error)
	GetAdminByUserID(ctx context.Context, userID int64) (entities.Admin, error)
	UpdateAdmin(ctx context.Context, admin entities.Admin) (entities.Admin, error)
	DeleteAdmin(ctx context.Context, userID int64) error
}
