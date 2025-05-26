package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// Create User and Customer in a single transaction
func (store *SQLStore) CreateCustomerTx(ctx context.Context, arg CreateUserParams) (User, error) {
	var user User

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		user, err = q.CreateUser(ctx, CreateUserParams{
			FirstName:      arg.FirstName,
			LastName:       arg.LastName,
			HashedPassword: arg.HashedPassword,
			Email:          arg.Email,
			Role:           UserRoleCustomer,
		})
		if err != nil {
			return err
		}
		_, err = q.CreateCustomer(ctx, CreateCustomerParams{
			UserID:               user.UserID,
			PhoneNumber:          pgtype.Text{String: "", Valid: false},
			Gender:               "other",
			DateOfBirth:          pgtype.Date{Time: time.Time{}, Valid: false},
			PassportNumber:       pgtype.Text{String: "", Valid: false},
			IdentificationNumber: pgtype.Text{String: "", Valid: false},
			Address:              pgtype.Text{String: "", Valid: false},
		})
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return User{}, err
	}

	return user, nil
}
