package postgresql

import (
	"context"

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
	// config, err := config.LoadConfig(".")
	// if err != nil {
	// 	log.Fatal("cannot load config:", err)
	// }
	// accessToken, accessPayload, err := r.tokenMaker.CreateToken(user.FirstName+user.LastName, user.Role, config.AccessTokenDuration, token.TokenTypeAccessToken)
	user, err := r.store.GetUser(ctx, userID)
	if err != nil {
		return entities.User{}, err
	}

	return entities.User{
		UserID:         user.UserID,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		HashedPassword: user.HashedPassword,
		Role:           entities.UserRole(user.Role),
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
