package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/spaghetti-lover/qairlines/db/sqlc"
	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
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
		FirstName:      &arg.FirstName,
		LastName:       &arg.LastName,
		HashedPassword: hashedPassword,
		Email:          arg.Email,
		Role:           db.UserRoleCustomer,
	})

	if err != nil {
		return entities.User{}, err
	}
	return entities.User{
		UserID:    user.UserID,
		FirstName: *user.FirstName,
		LastName:  *user.LastName,
		Email:     user.Email,
		HashedPwd: user.HashedPassword,
		Role:      entities.UserRole(entities.RoleCustomer),
	}, nil
}

func (r *CustomerRepositoryPostgres) CreateCustomer(ctx context.Context, arg entities.CreateCustomerParams) (entities.Customer, error) {
	customers, err := r.store.CreateCustomer(ctx, db.CreateCustomerParams{
		UserID:               arg.UserID,
		PhoneNumber:          &arg.PhoneNumber,
		Gender:               db.GenderType(arg.Gender),
		DateOfBirth:          pgtype.Date{Time: arg.DateOfBirth, Valid: true},
		PassportNumber:       &arg.PassportNumber,
		IdentificationNumber: &arg.IdentificationNumber,
		Address:              &arg.Address,
		LoyaltyPoints:        &arg.LoyaltyPoints,
	})
	if err != nil {
		return entities.Customer{}, err
	}
	return entities.Customer{
		UserID:               customers.UserID,
		PhoneNumber:          *customers.PhoneNumber,
		Gender:               entities.CustomerGender(customers.Gender),
		DateOfBirth:          customers.DateOfBirth.Time,
		PassportNumber:       *customers.PassportNumber,
		IdentificationNumber: *customers.IdentificationNumber,
		Address:              *customers.Address,
		LoyaltyPoints:        *customers.LoyaltyPoints,
	}, nil
}

func (r *CustomerRepositoryPostgres) UpdateCustomer(ctx context.Context, customer entities.Customer, user entities.User) (entities.Customer, entities.User, error) {
	err := r.store.UpdateCustomerTx(ctx, db.UpdateCustomerTxParams{
		UserID:               customer.UserID,
		FirstName:            user.FirstName,
		LastName:             user.LastName,
		PhoneNumber:          customer.PhoneNumber,
		Gender:               entities.GenderType(customer.Gender),
		DateOfBirth:          pgtype.Date{Time: customer.DateOfBirth, Valid: true},
		Address:              customer.Address,
		PassportNumber:       customer.PassportNumber,
		IdentificationNumber: customer.IdentificationNumber,
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
		}, entities.User{
			UserID:    user.UserID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}, nil
}

func (r *CustomerRepositoryPostgres) GetAllCustomers(ctx context.Context) ([]entities.Customer, error) {
	rows, err := r.store.GetAllCustomers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all customers: %w", err)
	}

	var customers []entities.Customer
	for _, row := range rows {
		user, err := r.store.GetUser(ctx, row.Uid)
		if err != nil {
			return nil, err
		}
		customers = append(customers, entities.Customer{
			UserID: row.Uid,
			User: entities.User{
				FirstName: *user.FirstName,
				LastName:  *user.LastName,
				Email:     user.Email,
			},
			DateOfBirth:          row.DateOfBirth.Time,
			Gender:               entities.CustomerGender(row.Gender),
			LoyaltyPoints:        *row.LoyaltyPoints,
			Address:              *row.Address,
			PassportNumber:       *row.PassportNumber,
			IdentificationNumber: *row.IdentificationNumber,
		})
	}

	return customers, nil
}

func (r *CustomerRepositoryPostgres) DeleteCustomerByID(ctx context.Context, customerID int64) error {
	customerID, err := r.store.DeleteCustomerByID(ctx, customerID)
	if err == sql.ErrNoRows {
		return adapters.ErrCustomerNotFound
	}
	if customerID == 0 {
		return adapters.ErrCustomerNotFound
	}
	if err != nil {
		return fmt.Errorf("failed to delete customer: %w", err)
	}

	return nil
}

func (r *CustomerRepositoryPostgres) GetCustomerByUID(ctx context.Context, uid int64) (*entities.Customer, error) {
	row, err := r.store.GetCustomerByID(ctx, uid)
	if err == sql.ErrNoRows {
		return nil, adapters.ErrCustomerNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get customer by UID: %w", err)
	}

	return &entities.Customer{
		UserID: row.Uid,
		User: entities.User{
			FirstName: *row.FirstName,
			LastName:  *row.LastName,
			Email:     row.Email,
		},
		PhoneNumber:          *row.PhoneNumber,
		DateOfBirth:          row.DateOfBirth.Time,
		Gender:               entities.CustomerGender(row.Gender),
		IdentificationNumber: *row.IdentificationNumber,
		PassportNumber:       *row.PassportNumber,
		Address:              *row.Address,
		LoyaltyPoints:        *row.LoyaltyPoints,
	}, nil
}

func (r *CustomerRepositoryPostgres) GetBookingHistoryByUID(ctx context.Context, uid int64) ([]string, error) {
	rows, err := r.store.GetBookingHistoryByUID(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("failed to get booking history by UID: %w", err)
	}

	var bookingHistory []string
	for _, row := range rows {
		bookingHistory = append(bookingHistory, strconv.FormatInt(row, 10))
	}

	return bookingHistory, nil
}
