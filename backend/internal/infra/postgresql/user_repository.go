package postgresql

import (
	"context"

	db "github.com/spaghetti-lover/qairlines/db/sqlc"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type UserRepositoryPostgres struct {
	store db.Store
}

// NewUserRepositoryPostgres creates a new UserRepositoryPostgres instance through using dependency injection with store is dependency that is injected
// into the UserRepositoryPostgres struct. This allows for better separation of concerns and makes it easier to test the code.
func NewUserRepositoryPostgres(store *db.Store) *UserRepositoryPostgres {
	return &UserRepositoryPostgres{
		store: *store,
	}
}

func (r *UserRepositoryPostgres) GetAllUser(ctx context.Context) ([]entities.User, error) {
	users, err := r.store.GetAllUser(ctx)
	if err != nil {
		return nil, err
	}
	usersList := make([]entities.User, len(users))
	for i, user := range users {
		usersList[i] = entities.User{
			UserID:         user.UserID,
			Username:       user.Username,
			HashedPassword: user.HashedPassword,
			Role:           entities.UserRole(user.Role),
		}
	}
	return usersList, nil
}

func (r *UserRepositoryPostgres) GetUser(ctx context.Context, userID int64) (entities.User, error) {
	user, err := r.store.GetUser(ctx, userID)
	if err != nil {
		return entities.User{}, err
	}
	return entities.User{
		UserID:         user.UserID,
		Username:       user.Username,
		HashedPassword: user.HashedPassword,
		Role:           entities.UserRole(user.Role),
	}, nil
}

func (r *UserRepositoryPostgres) CreateUser(ctx context.Context, arg entities.CreateUserParams) (entities.User, error) {
	user, err := r.store.CreateUser(ctx, db.CreateUserParams{
		Username:       arg.Username,
		HashedPassword: arg.Password,
		Role:           string(arg.Role),
	})
	if err != nil {
		return entities.User{}, err
	}
	return entities.User{
		UserID:         user.UserID,
		Username:       user.Username,
		HashedPassword: user.HashedPassword,
		Role:           entities.UserRole(user.Role),
	}, nil
}

func (r *UserRepositoryPostgres) DeleteUser(ctx context.Context, userID int64) error {
	err := r.store.DeleteUser(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryPostgres) ListUsers(ctx context.Context, arg entities.ListUsersParams) ([]entities.User, error) {
	users, err := r.store.ListUsers(ctx, db.ListUsersParams{
		Limit:  arg.Limit,
		Offset: arg.Offset,
	})
	if err != nil {
		return nil, err
	}
	usersList := make([]entities.User, len(users))
	for i, user := range users {
		usersList[i] = entities.User{
			UserID:         user.UserID,
			Username:       user.Username,
			HashedPassword: user.HashedPassword,
			Role:           entities.UserRole(user.Role),
		}
	}
	return usersList, nil
}
