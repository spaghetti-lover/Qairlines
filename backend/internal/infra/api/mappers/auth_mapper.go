package mappers

import (
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/auth"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
)

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type UserResponse struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}

func LoginOutputToResponse(output auth.LoginOutput) LoginResponse {
	return LoginResponse{
		Token: output.Token,
		User: UserResponse{
			ID:    output.User.UserID,
			Email: output.User.Email,
			Name:  output.User.FirstName + " " + output.User.LastName,
			Role:  string(output.User.Role),
		},
	}
}

func ChangePasswordRequestToInput(req dto.ChangePasswordRequest) auth.ChangePasswordInput {
	return auth.ChangePasswordInput{
		Email:       req.Email,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	}
}

func ChangePasswordResponseFromResult() dto.ChangePasswordResponse {
	return dto.ChangePasswordResponse{
		Message: "Password changed successfully.",
	}
}
