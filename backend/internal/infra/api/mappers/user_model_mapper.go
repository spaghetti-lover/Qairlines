package mappers

import (
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
)

func UserGetOutputToResponse(output usecases.UserGetOutput) dto.UserGetResponse {
	return dto.UserGetResponse{
		UserID:   output.UserID,
		Username: output.Name,
		Password: output.Password,
		Role:     output.Role,
	}
}
