package customer

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
)

type IGetAllCustomersUseCase interface {
	Execute(ctx context.Context) ([]dto.CustomerResponse, error)
}

type GetAllCustomersUseCase struct {
	customerRepository adapters.ICustomerRepository
}

func NewGetAllCustomersUseCase(customerRepository adapters.ICustomerRepository) IGetAllCustomersUseCase {
	return &GetAllCustomersUseCase{
		customerRepository: customerRepository,
	}
}

func (u *GetAllCustomersUseCase) Execute(ctx context.Context) ([]dto.CustomerResponse, error) {
	// Lấy danh sách khách hàng từ repository
	customers, err := u.customerRepository.GetAllCustomers(ctx)
	if err != nil {
		return nil, err
	}

	// Sử dụng mapper để chuyển đổi danh sách entity sang DTO
	return mappers.ToCustomerResponses(customers), nil
}
