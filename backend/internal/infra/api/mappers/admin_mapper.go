package mappers

import (
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/admin"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
)

func CreateAdminInputToRequest(input dto.CreateAdminRequest) entities.CreateAdminParams {
	return entities.CreateAdminParams{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
	}
}
func CreateAdminGetOutputToResponse(output entities.User) map[string]interface{} {
	return map[string]interface{}{
		"message": "Tài khoản quản trị viên đã được tạo thành công.",
		"admin": map[string]interface{}{
			"id":        output.UserID,
			"firstName": output.FirstName,
			"lastName":  output.LastName,
			"email":     output.Email,
			"createdAt": map[string]int64{
				"seconds": output.CreatedAt.Unix(),
			},
		},
	}

}

func AdminEntityToResponse(admin entities.Admin) dto.AdminResponse {
	response := dto.AdminResponse{
		ID:        admin.UserID,
		FirstName: admin.FirstName,
		LastName:  admin.LastName,
		Email:     admin.Email,
	}
	response.CreatedAt.Seconds = admin.CreatedAt.Unix()
	return response
}

func AdminsEntitiesToResponse(admins []entities.Admin) []dto.AdminResponse {
	responses := make([]dto.AdminResponse, len(admins))
	for i, admin := range admins {
		responses[i] = AdminEntityToResponse(admin)
	}
	return responses
}

func CurrentAdminEntityToResponse(admin entities.Admin) dto.GetCurrentAdminResponse {
	response := dto.GetCurrentAdminResponse{
		Message: "Admin retrieved successfully.",
	}
	response.Data.UID = admin.UserID
	response.Data.FirstName = admin.FirstName
	response.Data.LastName = admin.LastName
	response.Data.Email = admin.Email
	return response
}

func AdminUpdateRequestToInput(req dto.AdminUpdateRequest, userID int64) admin.UpdateAdminInput {
	return admin.UpdateAdminInput{
		UserID:    userID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}
}

func AdminUpdateEntityToResponse(admin entities.Admin) dto.AdminUpdateResponse {
	response := dto.AdminUpdateResponse{
		Message: "Admin updated successfully.",
	}
	response.Data.UID = admin.UserID
	response.Data.FirstName = admin.FirstName
	response.Data.LastName = admin.LastName
	response.Data.Email = admin.Email
	return response
}
