package user

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type IUserUpdateUseCase interface {
	Execute(ctx context.Context, id int64, user entities.User) (entities.User, error)
}

type UserUpdateUseCase struct {
	userRepository adapters.IUserRepository
}

func NewUserUpdateUseCase(userRepository adapters.IUserRepository) IUserUpdateUseCase {
	return &UserUpdateUseCase{
		userRepository: userRepository,
	}
}
func (u *UserUpdateUseCase) Execute(ctx context.Context, id int64, user entities.User) (entities.User, error) {
	user.UserID = id
	user, err := u.userRepository.UpdateUser(ctx, entities.UpdateUserParams{
		UserID:    id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	})
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}
