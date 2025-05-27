package adapters

import (
	"context"
	"errors"

	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

var (
	ErrUnauthorized     = errors.New("authentication failed. Admin privileges required")
	ErrCustomerNotFound = errors.New("customer not found")
)

type ICustomerRepository interface {
	CreateCustomerTx(ctx context.Context, arg entities.CreateUserParams) (entities.User, error)
	UpdateCustomer(ctx context.Context, costumer entities.Customer, user entities.User) (entities.Customer, entities.User, error)
	GetAllCustomers(ctx context.Context) ([]entities.Customer, error)
	DeleteCustomerByID(ctx context.Context, customerID int64) error
}
