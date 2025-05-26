package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/pkg/utils"
)

// Define errors
var (
	ErrOldPasswordIncorrect     = errors.New("old password is incorrect")
	ErrPasswordValidationFailed = errors.New("new password does not meet the required criteria")
	ErrUserNotFound             = errors.New("user not found")
)

// Định nghĩa input cho use case đổi mật khẩu
type ChangePasswordInput struct {
	Email       string
	OldPassword string
	NewPassword string
}

// Interface cho use case đổi mật khẩu
type IChangePasswordUseCase interface {
	Execute(ctx context.Context, input ChangePasswordInput) error
}

// Implement use case
type ChangePasswordUseCase struct {
	userRepository adapters.IUserRepository
}

// Constructor
func NewChangePasswordUseCase(userRepository adapters.IUserRepository) IChangePasswordUseCase {
	return &ChangePasswordUseCase{
		userRepository: userRepository,
	}
}

// Execute thực hiện việc đổi mật khẩu
func (u *ChangePasswordUseCase) Execute(ctx context.Context, input ChangePasswordInput) error {
	// 1. Kiểm tra người dùng tồn tại
	if input.Email == "" {
		return fmt.Errorf("%w: email cannot be empty", ErrUserNotFound)
	}
	if input.OldPassword == "" {
		return fmt.Errorf("%w: old password cannot be empty", ErrOldPasswordIncorrect)
	}
	if input.NewPassword == "" {
		return fmt.Errorf("%w: new password cannot be empty", ErrPasswordValidationFailed)
	}
	user, err := u.userRepository.GetUserByEmail(ctx, input.Email)
	if err != nil {
		// Xử lý lỗi "no rows"
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("%w: email %s", ErrUserNotFound, input.Email)
		}
		return fmt.Errorf("error fetching user: %w", err)
	}

	// 2. Kiểm tra mật khẩu cũ có đúng không
	err = utils.CheckPassword(input.OldPassword, user.HashedPwd)
	if err != nil {
		return ErrOldPasswordIncorrect
	}

	// 3. Hash mật khẩu mới
	hashedPassword, err := utils.HashPassword(input.NewPassword)
	if err != nil {
		return err
	}

	// 4. Cập nhật mật khẩu
	return u.userRepository.UpdatePassword(ctx, input.Email, hashedPassword)
}
