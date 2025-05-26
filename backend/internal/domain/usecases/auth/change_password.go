package auth

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	appErrors "github.com/spaghetti-lover/qairlines/pkg/errors"
	"github.com/spaghetti-lover/qairlines/pkg/utils"
)

type ChangePasswordInput struct {
	UserID      int64
	OldPassword string
	NewPassword string
}

type IChangePasswordUseCase interface {
	Execute(ctx context.Context, input ChangePasswordInput) error
}

type ChangePasswordUseCase struct {
	userRepository adapters.IUserRepository
}

func NewChangePasswordUseCase(userRepository adapters.IUserRepository) IChangePasswordUseCase {
	return &ChangePasswordUseCase{userRepository: userRepository}
}

func (u *ChangePasswordUseCase) Execute(ctx context.Context, input ChangePasswordInput) error {
	// Lấy thông tin người dùng
	user, err := u.userRepository.GetUser(ctx, input.UserID)
	if err != nil {
		return &appErrors.AppError{Message: "User not found."}
	}

	// Kiểm tra mật khẩu cũ
	err = utils.CheckPassword(input.OldPassword, user.HashedPwd)
	if err != nil {
		return &appErrors.AppError{Message: "Old password is incorrect."}
	}

	// Cập nhật mật khẩu mới
	err = u.userRepository.UpdatePassword(ctx, input.UserID, input.NewPassword)
	if err != nil {
		return &appErrors.AppError{Message: "Failed to update password."}
	}

	return nil
}
