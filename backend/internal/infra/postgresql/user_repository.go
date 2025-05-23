package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/spaghetti-lover/qairlines/db/sqlc"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
	"github.com/spaghetti-lover/qairlines/pkg/token"
	"github.com/spaghetti-lover/qairlines/pkg/utils"
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
			UserID:         user.UserID,
			FirstName:      user.FirstName,
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
		UserID:               user.UserID,
		FirstName:            user.FirstName,
		LastName:             user.LastName,
		HashedPassword:       user.HashedPassword,
		PhoneNumber:          user.PhoneNumber.String,
		Gender:               user.Gender.String,
		Address:              user.Address.String,
		PassportNumber:       user.PassportNumber.String,
		IdentificationNumber: user.IdentificationNumber.String,
		Role:                 entities.UserRole(user.Role),
		Email:                user.Email,
		LoyaltyPoints:        user.LoyaltyPoints.Int64,
		CreatedAt:            user.CreatedAt,
		UpdatedAt:            user.UpdatedAt,
	}, nil
}

func (r *UserRepositoryPostgres) CreateUser(ctx context.Context, arg entities.CreateUserParams) (entities.User, error) {
	hashedPassword, err := utils.HashPassword(arg.Password)
	if err != nil {
		return entities.User{}, err
	}
	user, err := r.store.CreateUser(ctx, db.CreateUserParams{
		FirstName:      arg.FirstName,
		LastName:       arg.LastName,
		HashedPassword: hashedPassword,
		Email:          arg.Email,
	})

	if err != nil {
		return entities.User{}, err
	}
	return entities.User{
		UserID:         user.UserID,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
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
			FirstName:      user.FirstName,
			LastName:       user.LastName,
			HashedPassword: user.HashedPassword,
			Role:           entities.UserRole(user.Role),
		}
	}
	return usersList, nil
}

func (r *UserRepositoryPostgres) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	user, err := r.store.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &entities.User{
		UserID:         user.UserID,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		Role:           entities.UserRole(user.Role),
	}, nil
}

func (r *UserRepositoryPostgres) UpdatePassword(ctx context.Context, userID int64, newPassword string) error {
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	err = r.store.UpdatePassword(ctx, db.UpdatePasswordParams{
		UserID:         userID,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryPostgres) UpdateUser(ctx context.Context, user entities.User) error {
	err := r.store.UpdateUser(ctx, db.UpdateUserParams{
		UserID:               user.UserID,
		FirstName:            user.FirstName,
		LastName:             user.LastName,
		PhoneNumber:          pgtype.Text{String: user.PhoneNumber, Valid: true},
		Gender:               pgtype.Text{String: user.Gender, Valid: true},
		Address:              pgtype.Text{String: user.Address, Valid: true},
		PassportNumber:       pgtype.Text{String: user.PassportNumber, Valid: true},
		IdentificationNumber: pgtype.Text{String: user.IdentificationNumber, Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}
