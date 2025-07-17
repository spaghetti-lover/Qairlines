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
		FirstName:      pgtype.Text{String: arg.FirstName, Valid: true},
		LastName:       pgtype.Text{String: arg.LastName, Valid: true},
		HashedPassword: hashedPassword,
		Email:          arg.Email,
		Role:           db.UserRoleCustomer,
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
		DateOfBirth:          arg.DateOfBirth,
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
		DateOfBirth:          customers.DateOfBirth,
		PassportNumber:       customers.PassportNumber.String,
		IdentificationNumber: customers.IdentificationNumber.String,
		Address:              customers.Address.String,
		LoyaltyPoints:        customers.LoyaltyPoints.Int32,
	}, nil
}

func (r *CustomerRepositoryPostgres) UpdateCustomer(ctx context.Context, customer entities.Customer, user entities.User) (entities.Customer, entities.User, error) {
	err := r.store.UpdateCustomerTx(ctx, db.UpdateCustomerTxParams{
		UserID:               customer.UserID,
		FirstName:            user.FirstName,
		LastName:             user.LastName,
		PhoneNumber:          customer.PhoneNumber,
		Gender:               entities.GenderType(customer.Gender),
		DateOfBirth:          customer.DateOfBirth,
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

func (r *CustomerRepositoryPostgres) ListCustomers(ctx context.Context, offset int, limit int) ([]entities.Customer, error) {
	rows, err := r.store.ListCustomers(ctx, db.ListCustomersParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list customers: %w", err)
	}

	var customers []entities.Customer
	for _, row := range rows {
		user, err := r.store.GetUser(ctx, row.UserID)
		if err != nil {
			return nil, err
		}
		customers = append(customers, entities.Customer{
			UserID: row.UserID,
			User: entities.User{
				FirstName: user.FirstName.String,
				LastName:  user.LastName.String,
				Email:     user.Email,
			},
			DateOfBirth:          row.DateOfBirth,
			Gender:               entities.CustomerGender(row.Gender),
			LoyaltyPoints:        row.LoyaltyPoints.Int32,
			Address:              row.Address.String,
			PassportNumber:       row.PassportNumber.String,
			IdentificationNumber: row.IdentificationNumber.String,
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
			FirstName: row.FirstName.String,
			LastName:  row.LastName.String,
			Email:     row.Email,
		},
		PhoneNumber:          row.PhoneNumber.String,
		DateOfBirth:          row.DateOfBirth,
		Gender:               entities.CustomerGender(row.Gender),
		IdentificationNumber: row.IdentificationNumber.String,
		PassportNumber:       row.PassportNumber.String,
		Address:              row.Address.String,
		LoyaltyPoints:        row.LoyaltyPoints.Int32,
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
