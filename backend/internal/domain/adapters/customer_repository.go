package adapters

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type ICustomerRepository interface {
	CreateCustomerTx(ctx context.Context, arg entities.CreateUserParams) (entities.User, error)
	UpdateCustomer(ctx context.Context, costumer entities.Customer, user entities.User) (entities.Customer, entities.User, error)
}
