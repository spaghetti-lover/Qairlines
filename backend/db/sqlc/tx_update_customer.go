package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type UpdateCustomerTxParams struct {
	UserID               int64               `json:"user_id"`
	FirstName            string              `json:"first_name"`
	LastName             string              `json:"last_name"`
	PhoneNumber          pgtype.Text         `json:"phone_number"`
	Gender               entities.GenderType `json:"gender"`
	Address              pgtype.Text         `json:"address"`
	DateOfBirth          pgtype.Date         `json:"date_of_birth"`
	PassportNumber       pgtype.Text         `json:"passport_number"`
	IdentificationNumber pgtype.Text         `json:"identification_number"`
}

func (store *SQLStore) UpdateCustomerTx(ctx context.Context, arg UpdateCustomerTxParams) error {
	return store.execTx(ctx, func(q *Queries) error {
		// Update customer
		err := store.UpdateCustomer(ctx, UpdateCustomerParams{
			UserID:               arg.UserID,
			PhoneNumber:          arg.PhoneNumber,
			Gender:               GenderType(arg.Gender),
			DateOfBirth:          arg.DateOfBirth,
			Address:              arg.Address,
			PassportNumber:       arg.PassportNumber,
			IdentificationNumber: arg.IdentificationNumber,
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
