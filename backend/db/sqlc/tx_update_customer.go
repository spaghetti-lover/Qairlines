package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type UpdateCustomerTxParams struct {
	UserID               int64               `json:"user_id"`
	FirstName            string              `json:"first_name"`
	LastName             string              `json:"last_name"`
	PhoneNumber          string              `json:"phone_number"`
	Gender               entities.GenderType `json:"gender"`
	Address              string              `json:"address"`
	DateOfBirth          time.Time           `json:"date_of_birth"`
	PassportNumber       string              `json:"passport_number"`
	IdentificationNumber string              `json:"identification_number"`
}

func (store *SQLStore) UpdateCustomerTx(ctx context.Context, arg UpdateCustomerTxParams) error {
	return store.execTx(ctx, func(q *Queries) error {
		// Update customer
		err := store.UpdateCustomer(ctx, UpdateCustomerParams{
			UserID:               arg.UserID,
			PhoneNumber:          pgtype.Text{String: arg.PhoneNumber, Valid: true},
			Gender:               GenderType(arg.Gender),
			DateOfBirth:          arg.DateOfBirth,
			Address:              pgtype.Text{String: arg.Address, Valid: true},
			PassportNumber:       pgtype.Text{String: arg.PassportNumber, Valid: true},
			IdentificationNumber: pgtype.Text{String: arg.IdentificationNumber, Valid: true},
		})
		if err != nil {
			return err
		}
		// Update user
		user := entities.User{
			UserID:    arg.UserID,
			FirstName: arg.FirstName,
			LastName:  arg.LastName,
		}
		err = store.UpdateUser(ctx, UpdateUserParams{
			UserID:    user.UserID,
			FirstName: pgtype.Text{String: user.FirstName, Valid: true},
			LastName:  pgtype.Text{String: user.LastName, Valid: true},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
