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

func UserUpdateInputToRequest(input dto.UserUpdateRequest) entities.User {
	return entities.User{
		FirstName:            input.FirstName,
		LastName:             input.LastName,
		PhoneNumber:          input.PhoneNumber,
		Gender:               input.Gender,
		Address:              input.Address,
		PassportNumber:       input.PassportNumber,
		IdentificationNumber: input.IdentificationNumber,
	}
}

func UserUpdateOutputToResponse(output entities.User) map[string]interface{} {
	return map[string]interface{}{
		"message": "Customer information updated successfully.",
	}
}

func UserGetOutputToResponseByToken(output entities.User) map[string]interface{} {
	return map[string]interface{}{
		"data": map[string]interface{}{
			"uid":                  output.UserID,
			"role":                 output.Role,
			"phoneNumber":          output.PhoneNumber,
			"dateOfBirth":          output.DateOfBirth,
			"firstName":            output.FirstName,
			"lastName":             output.LastName,
			"gender":               output.Gender,
			"email":                output.Email,
			"identificationNumber": output.IdentificationNumber,
			"passportNumber":       output.PassportNumber,
			"address":              output.Address,
			"loyaltyPoints":        output.LoyaltyPoints,
			"bookingHistory": []string{"BK174758329151929",
				"BK174784030948893",
				"BK174784036747874"},
			"createdAt": map[string]interface{}{
				"seconds":     output.CreatedAt.Unix(),
				"nanoseconds": output.CreatedAt.Nanosecond(),
			},
			"updatedAt": map[string]interface{}{
				"seconds":     output.UpdatedAt.Unix(),
				"nanoseconds": output.UpdatedAt.Nanosecond(),
			},
		},
	}
}
