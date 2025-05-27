package customer

import (
	"context"
	"errors"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
)

type IDeleteCustomerUseCase interface {
	Execute(ctx context.Context, customerID int64) error
}

type DeleteCustomerUseCase struct {
	customerRepository adapters.ICustomerRepository
}

func NewDeleteCustomerUseCase(customerRepository adapters.ICustomerRepository) IDeleteCustomerUseCase {
	return &DeleteCustomerUseCase{
		customerRepository: customerRepository,
	}
}

func (u *DeleteCustomerUseCase) Execute(ctx context.Context, customerID int64) error {
	// Xóa khách hàng trong repository
	err := u.customerRepository.DeleteCustomerByID(ctx, customerID)
	if err != nil {
		if errors.Is(err, adapters.ErrCustomerNotFound) {
			return adapters.ErrCustomerNotFound
		}
		return err
	}

	return nil
}
