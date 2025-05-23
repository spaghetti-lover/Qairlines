package user

import (
	"context"
	"errors"
	"time"

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
	return &UserUpdateUseCase{userRepository: userRepository}
}

func (u *UserUpdateUseCase) Execute(ctx context.Context, id int64, user entities.User) (entities.User, error) {
	// Kiểm tra xem user có tồn tại không
	existingUser, err := u.userRepository.GetUser(ctx, id)
	if err != nil {
		return entities.User{}, errors.New("user not found")
	}

	// Cập nhật thông tin user
	existingUser.FirstName = user.FirstName
	existingUser.LastName = user.LastName
	existingUser.PhoneNumber = user.PhoneNumber
	existingUser.Gender = user.Gender
	existingUser.Address = user.Address
	existingUser.PassportNumber = user.PassportNumber
	existingUser.IdentificationNumber = user.IdentificationNumber
	existingUser.UpdatedAt = time.Now()

	// Lưu thông tin cập nhật
	err = u.userRepository.UpdateUser(ctx, existingUser)
	if err != nil {
		return entities.User{}, errors.New("failed to update user")
	}

	return existingUser, nil
}
