package mappers

import (
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
)

func UserCreateInputToRequest(input dto.UserCreateRequest) entities.CreateUserParams {
	return entities.CreateUserParams{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Password:  input.Password,
		Email:     input.Email,
	}
}

func UserGetOutputToResponse(output entities.User) dto.UserGetResponse {
	return dto.UserGetResponse{
		UserID:    output.UserID,
		FirstName: output.FirstName,
		LastName:  output.LastName,
		Email:     output.Email,
	}
}

func UserGetListOutputToResponse(output []entities.User) []dto.UserGetResponse {
	responses := make([]dto.UserGetResponse, len(output))
	for i, user := range output {
		responses[i] = UserGetOutputToResponse(user)
	}
	return responses
}
