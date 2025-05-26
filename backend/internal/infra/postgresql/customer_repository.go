package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/spaghetti-lover/qairlines/db/sqlc"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
	"github.com/spaghetti-lover/qairlines/pkg/token"
	"github.com/spaghetti-lover/qairlines/pkg/utils"
)

type CustomerRepositoryPostgres struct {
	store      db.Store
	tokenMaker token.Maker
}

// NewUserRepositoryPostgres creates a new UserRepositoryPostgres instance through using dependency injection with store is dependency that is injected
// into the UserRepositoryPostgres struct. This allows for better separation of concerns and makes it easier to test the code.
func NewCustomerRepositoryPostgres(store *db.Store, tokenMaker token.Maker) *CustomerRepositoryPostgres {
	return &CustomerRepositoryPostgres{
		store:      *store,
		tokenMaker: tokenMaker,
	}
}

func (r *CustomerRepositoryPostgres) CreateCustomerTx(ctx context.Context, arg entities.CreateUserParams) (entities.User, error) {
	hashedPassword, err := utils.HashPassword(arg.Password)
	if err != nil {
		return entities.User{}, err
	}
	user, err := r.store.CreateCustomerTx(ctx, db.CreateUserParams{
		FirstName:      pgtype.Text{String: arg.FirstName, Valid: true},
		LastName:       pgtype.Text{String: arg.LastName, Valid: true},
		HashedPassword: hashedPassword,
		Email:          arg.Email,
	})

	if err != nil {
		return entities.User{}, err
	}
	return entities.User{
		UserID:    user.UserID,
		FirstName: user.FirstName.String,
		LastName:  user.LastName.String,
		Email:     user.Email,
		HashedPwd: user.HashedPassword,
		Role:      entities.UserRole(entities.RoleCustomer),
	}, nil
}

func (r *CustomerRepositoryPostgres) CreateCustomer(ctx context.Context, arg entities.CreateCustomerParams) (entities.Customer, error) {
	customers, err := r.store.CreateCustomer(ctx, db.CreateCustomerParams{
		UserID:               arg.UserID,
		PhoneNumber:          pgtype.Text{String: arg.PhoneNumber, Valid: true},
		Gender:               db.GenderType(arg.Gender),
		DateOfBirth:          pgtype.Date{Time: arg.DateOfBirth, Valid: true},
		PassportNumber:       pgtype.Text{String: arg.PassportNumber, Valid: true},
		IdentificationNumber: pgtype.Text{String: arg.IdentificationNumber, Valid: true},
		Address:              pgtype.Text{String: arg.Address, Valid: true},
		LoyaltyPoints:        pgtype.Int4{Int32: arg.LoyaltyPoints, Valid: true},
	})
	if err != nil {
		return entities.Customer{}, err
	}
	return entities.Customer{
		UserID:               customers.UserID,
		PhoneNumber:          customers.PhoneNumber.String,
		Gender:               entities.CustomerGender(customers.Gender),
		DateOfBirth:          customers.DateOfBirth.Time,
		PassportNumber:       customers.PassportNumber.String,
		IdentificationNumber: customers.IdentificationNumber.String,
		Address:              customers.Address.String,
		LoyaltyPoints:        customers.LoyaltyPoints.Int32,
		CreatedAt:            customers.CreatedAt,
		UpdatedAt:            customers.UpdatedAt,
	}, nil
}

func (r *CustomerRepositoryPostgres) UpdateCustomer(ctx context.Context, customer entities.Customer, user entities.User) (entities.Customer, entities.User, error) {
	err := r.store.UpdateCustomerTx(ctx, db.UpdateCustomerTxParams{
		UserID:               customer.UserID,
		FirstName:            user.FirstName,
		LastName:             user.LastName,
		PhoneNumber:          pgtype.Text{String: customer.PhoneNumber, Valid: true},
		Gender:               entities.GenderType(customer.Gender),
		DateOfBirth:          pgtype.Date{Time: customer.DateOfBirth, Valid: true},
		Address:              pgtype.Text{String: customer.Address, Valid: true},
		PassportNumber:       pgtype.Text{String: customer.PassportNumber, Valid: true},
		IdentificationNumber: pgtype.Text{String: customer.IdentificationNumber, Valid: true},
	})

	if err != nil {
		return entities.Customer{}, entities.User{}, err
	}
	_, err = r.store.GetCustomer(ctx, customer.UserID)
	if err != nil {
		return entities.Customer{}, entities.User{}, err
	}
	return entities.Customer{
			UserID:               customer.UserID,
			PhoneNumber:          customer.PhoneNumber,
			Gender:               entities.CustomerGender(customer.Gender),
			DateOfBirth:          customer.DateOfBirth,
			PassportNumber:       customer.PassportNumber,
			IdentificationNumber: customer.IdentificationNumber,
			Address:              customer.Address,
			LoyaltyPoints:        customer.LoyaltyPoints,
			CreatedAt:            customer.CreatedAt,
			UpdatedAt:            customer.UpdatedAt,
		}, entities.User{
			UserID:    user.UserID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}, nil
}
