package customer

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type ICustomerUpdateUseCase interface {
	Execute(ctx context.Context, id int64, customer entities.Customer, user entities.User) (entities.Customer, entities.User, error)
}

type CustomerUpdateUseCase struct {
	customerRepository adapters.ICustomerRepository
}

func NewCustomerUpdateUseCase(customerRepository adapters.ICustomerRepository) ICustomerUpdateUseCase {
	return &CustomerUpdateUseCase{customerRepository: customerRepository}
}

func (u *CustomerUpdateUseCase) Execute(ctx context.Context, id int64, customer entities.Customer, user entities.User) (entities.Customer, entities.User, error) {
	customer.UserID = id
	user.UserID = id
	customer, user, err := u.customerRepository.UpdateCustomer(ctx, customer, user)
	if err != nil {
		return entities.Customer{}, entities.User{}, err
	}
	return customer, user, nil
}
