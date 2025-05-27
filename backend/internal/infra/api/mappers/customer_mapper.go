package mappers

import (
	"strconv"

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

func ToCustomerResponses(customers []entities.Customer) []dto.CustomerResponse {
	var responses []dto.CustomerResponse

	for _, customer := range customers {
		responses = append(responses, dto.CustomerResponse{
			UID:                  strconv.FormatInt(customer.UserID, 10),
			FirstName:            customer.User.FirstName,
			LastName:             customer.User.LastName,
			Email:                customer.User.Email,
			DateOfBirth:          dto.TimeSeconds{Seconds: customer.DateOfBirth.Unix()},
			Gender:               string(customer.Gender),
			LoyaltyPoints:        customer.LoyaltyPoints,
			CreatedAt:            dto.TimeSeconds{Seconds: customer.CreatedAt.Unix()},
			Address:              customer.Address,
			PassportNumber:       customer.PassportNumber,
			IdentificationNumber: customer.IdentificationNumber,
		})
	}
	return responses
}

func ToCustomerDetailsResponse(customer *entities.Customer, bookingHistory []string) *dto.CustomerDetailsResponse {
	return &dto.CustomerDetailsResponse{
		UID:                  strconv.FormatInt(customer.UserID, 10),
		Role:                 "customer",
		PhoneNumber:          customer.PhoneNumber,
		DateOfBirth:          customer.DateOfBirth.Format("2006-01-02T15:04:05.000Z"),
		FirstName:            customer.User.FirstName,
		LastName:             customer.User.LastName,
		Gender:               string(customer.Gender),
		Email:                customer.User.Email,
		IdentificationNumber: &customer.IdentificationNumber,
		PassportNumber:       customer.PassportNumber,
		Address:              customer.Address,
		LoyaltyPoints:        int(customer.LoyaltyPoints),
		BookingHistory:       bookingHistory,
		CreatedAt: dto.TimeWithNano{
			Seconds:     customer.CreatedAt.Unix(),
			Nanoseconds: int64(customer.CreatedAt.Nanosecond()),
		},
		UpdatedAt: dto.TimeWithNano{
			Seconds:     customer.UpdatedAt.Unix(),
			Nanoseconds: int64(customer.UpdatedAt.Nanosecond()),
		},
	}
}
