package mappers

import (
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
)

func CreateCustomerGetOutputToResponse(output entities.User) dto.CreateCustomerResponse {
	return dto.CreateCustomerResponse{
		Message: "Tài khoản đã được tạo thành công.",
		User: struct {
			ID        int64  `json:"id"`
			FirstName string `json:"firstName"`
			LastName  string `json:"lastName"`
			Email     string `json:"email"`
		}{
			ID:        output.UserID,
			FirstName: output.FirstName,
			LastName:  output.LastName,
			Email:     output.Email,
		},
	}
}

func CustomerUpdateResponse(customer entities.Customer, user entities.User) map[string]interface{} {
	return map[string]interface{}{
		"message": "Customer information updated successfully.",
		"data": dto.CustomerUpdateResponse{
			UID:                  customer.UserID,
			FirstName:            user.FirstName,
			LastName:             user.LastName,
			PhoneNumber:          customer.PhoneNumber,
			Gender:               entities.GenderType(customer.Gender),
			Address:              customer.Address,
			DateOfBirth:          dto.DateOfBirth{Seconds: customer.DateOfBirth.Unix()},
			Passport:             customer.PassportNumber,
			IdentificationNumber: customer.IdentificationNumber,
		},
	}
}
