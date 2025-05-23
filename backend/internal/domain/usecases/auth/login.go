package auth

import (
	"context"

	appErrors "github.com/spaghetti-lover/qairlines/pkg/errors"

	"github.com/spaghetti-lover/qairlines/config"
	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
	"github.com/spaghetti-lover/qairlines/pkg/token"
	"github.com/spaghetti-lover/qairlines/pkg/utils"
)

type LoginInput struct {
	Email    string
	Password string
}

type LoginOutput struct {
	Token string
	User  *entities.User
}

type ILoginUseCase interface {
	Execute(ctx context.Context, input LoginInput) (*LoginOutput, error)
}

type LoginUseCase struct {
	userRepository adapters.IUserRepository
	tokenMaker     token.Maker
}

func NewLoginUseCase(userRepository adapters.IUserRepository, tokenMaker token.Maker) ILoginUseCase {
	return &LoginUseCase{
		userRepository: userRepository,
		tokenMaker:     tokenMaker,
	}
}

func (u *LoginUseCase) Execute(ctx context.Context, input LoginInput) (*LoginOutput, error) {
	config, err := config.LoadConfig(".")
	if err != nil {
		return nil, err
	}
	// Get user info by email
	user, err := u.userRepository.GetUserByEmail(ctx, input.Email)
	if err != nil {
		message := utils.GetErrorMessage("ERR_USER_NOT_FOUND", "vi")
		return nil, &appErrors.AppError{Message: message}
	}
	if user == nil {
		message := utils.GetErrorMessage("ERR_INVALID_CREDENTIALS", "vi")
		return nil, &appErrors.AppError{Message: message}
	}

	// Verify password
	err = utils.CheckPassword(input.Password, user.HashedPassword)
	if err != nil {
		message := utils.GetErrorMessage("ERR_INVALID_CREDENTIALS", "vi")
		return nil, &appErrors.AppError{Message: message}
	}

	// Generate token
	accessToken, _, err := u.tokenMaker.CreateToken(user.UserID, string(user.Role), config.AccessTokenDuration, token.TokenTypeAccessToken)
	if err != nil {
		return nil, err
	}

	return &LoginOutput{
		Token: accessToken,
		User:  user,
	}, nil
}
