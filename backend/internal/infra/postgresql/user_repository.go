package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/spaghetti-lover/qairlines/db/sqlc"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
	"github.com/spaghetti-lover/qairlines/pkg/token"
)

type UserRepositoryPostgres struct {
	store      db.Store
	tokenMaker token.Maker
}

// NewUserRepositoryPostgres creates a new UserRepositoryPostgres instance through using dependency injection with store is dependency that is injected
// into the UserRepositoryPostgres struct. This allows for better separation of concerns and makes it easier to test the code.
func NewUserRepositoryPostgres(store *db.Store, tokenMaker token.Maker) *UserRepositoryPostgres {
	return &UserRepositoryPostgres{
		store:      *store,
		tokenMaker: tokenMaker,
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
			UserID:    user.UserID,
			FirstName: user.FirstName.String,
			LastName:  user.LastName.String,
			HashedPwd: user.HashedPassword,
			Role:      entities.UserRole(user.Role),
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
		UserID:    user.UserID,
		FirstName: user.FirstName.String,
		LastName:  user.LastName.String,
		HashedPwd: user.HashedPassword,
		Role:      entities.UserRole(user.Role),
		Email:     user.Email,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// func (r *UserRepositoryPostgres) DeleteUser(ctx context.Context, userID int64) error {
// 	err := r.store.DeleteUser(ctx, userID)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (r *UserRepositoryPostgres) ListUsers(ctx context.Context, arg entities.ListUsersParams) ([]entities.User, error) {
// 	users, err := r.store.ListUsers(ctx, db.ListUsersParams{
// 		Limit:  arg.Limit,
// 		Offset: arg.Offset,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	usersList := make([]entities.User, len(users))
// 	for i, user := range users {
// 		usersList[i] = entities.User{
// 			UserID:         user.UserID,
// 			FirstName:      user.FirstName,
// 			LastName:       user.LastName,
// 			HashedPassword: user.HashedPassword,
// 			Role:           entities.UserRole(user.Role),
// 		}
// 	}
// 	return usersList, nil
// }

func (r *UserRepositoryPostgres) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	user, err := r.store.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &entities.User{
		UserID:    user.UserID,
		FirstName: user.FirstName.String,
		LastName:  user.LastName.String,
		Email:     user.Email,
		HashedPwd: user.HashedPassword,
		Role:      entities.UserRole(user.Role),
	}, nil
}

func (r *UserRepositoryPostgres) UpdatePassword(ctx context.Context, email string, hashedPassword string) error {
	err := r.store.UpdateUserPassword(ctx, db.UpdateUserPasswordParams{
		Email:          email,
		HashedPassword: hashedPassword,
	})

	if err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}

	return nil
}

func (r *UserRepositoryPostgres) UpdateUser(ctx context.Context, arg entities.UpdateUserParams) (entities.User, error) {
	err := r.store.UpdateUser(ctx, db.UpdateUserParams{
		UserID:    arg.UserID,
		FirstName: pgtype.Text{String: arg.FirstName, Valid: true},
		LastName:  pgtype.Text{String: arg.LastName, Valid: true},
	})
	if err != nil {
		return entities.User{}, err
	}

	return entities.User{
		UserID:    arg.UserID,
		FirstName: arg.FirstName,
		LastName:  arg.LastName,
	}, nil
}
