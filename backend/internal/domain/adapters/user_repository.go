package adapters

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type IUserRepository interface {
	GetUser(ctx context.Context, userID int64) (entities.User, error)
	CreateUser(ctx context.Context, arg entities.CreateUserParams) (entities.User, error)
	DeleteUser(ctx context.Context, userID int64) error
	ListUsers(ctx context.Context, arg entities.ListUsersParams) ([]entities.User, error)
	GetAllUser(ctx context.Context) ([]entities.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
	UpdatePassword(ctx context.Context, userID int64, newPassword string) error
}
