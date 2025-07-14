package customer

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
)

type IListCustomersUseCase interface {
	Execute(ctx context.Context, page int, limit int) ([]dto.CustomerResponse, error)
}

type ListCustomersUseCase struct {
	customerRepository adapters.ICustomerRepository
}

func NewListCustomersUseCase(customerRepository adapters.ICustomerRepository) IListCustomersUseCase {
	return &ListCustomersUseCase{
		customerRepository: customerRepository,
	}
}

func (u *ListCustomersUseCase) Execute(ctx context.Context, page int, limit int) ([]dto.CustomerResponse, error) {
	start := (page - 1) * limit
	// Lấy danh sách khách hàng từ repository
	customers, err := u.customerRepository.ListCustomers(ctx, start, limit)
	if err != nil {
		return nil, err
	}

	// Sử dụng mapper để chuyển đổi danh sách entity sang DTO
	return mappers.ToCustomerResponses(customers), nil
}
