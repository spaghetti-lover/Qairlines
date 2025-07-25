// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: ticket_owner_snapshot.sql

package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTicketOwnerSnapshot = `-- name: CreateTicketOwnerSnapshot :one
INSERT INTO TicketOwnerSnapshots (
  ticket_id, first_name, last_name, phone_number, gender, date_of_birth,
  passport_number, identification_number, address
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING ticket_id, first_name, last_name, phone_number, gender, date_of_birth, passport_number, identification_number, address
`

type CreateTicketOwnerSnapshotParams struct {
	TicketID             int64       `json:"ticket_id"`
	FirstName            pgtype.Text `json:"first_name"`
	LastName             pgtype.Text `json:"last_name"`
	PhoneNumber          pgtype.Text `json:"phone_number"`
	Gender               GenderType  `json:"gender"`
	DateOfBirth          time.Time   `json:"date_of_birth"`
	PassportNumber       pgtype.Text `json:"passport_number"`
	IdentificationNumber pgtype.Text `json:"identification_number"`
	Address              pgtype.Text `json:"address"`
}

func (q *Queries) CreateTicketOwnerSnapshot(ctx context.Context, arg CreateTicketOwnerSnapshotParams) (Ticketownersnapshot, error) {
	row := q.db.QueryRow(ctx, createTicketOwnerSnapshot,
		arg.TicketID,
		arg.FirstName,
		arg.LastName,
		arg.PhoneNumber,
		arg.Gender,
		arg.DateOfBirth,
		arg.PassportNumber,
		arg.IdentificationNumber,
		arg.Address,
	)
	var i Ticketownersnapshot
	err := row.Scan(
		&i.TicketID,
		&i.FirstName,
		&i.LastName,
		&i.PhoneNumber,
		&i.Gender,
		&i.DateOfBirth,
		&i.PassportNumber,
		&i.IdentificationNumber,
		&i.Address,
	)
	return i, err
}

const getAllTicketOwnerSnapshots = `-- name: GetAllTicketOwnerSnapshots :many
SELECT ticket_id, first_name, last_name, phone_number, gender, date_of_birth, passport_number, identification_number, address FROM TicketOwnerSnapshots
`

func (q *Queries) GetAllTicketOwnerSnapshots(ctx context.Context) ([]Ticketownersnapshot, error) {
	rows, err := q.db.Query(ctx, getAllTicketOwnerSnapshots)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Ticketownersnapshot{}
	for rows.Next() {
		var i Ticketownersnapshot
		if err := rows.Scan(
			&i.TicketID,
			&i.FirstName,
			&i.LastName,
			&i.PhoneNumber,
			&i.Gender,
			&i.DateOfBirth,
			&i.PassportNumber,
			&i.IdentificationNumber,
			&i.Address,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTicketOwnerSnapshot = `-- name: GetTicketOwnerSnapshot :one
SELECT ticket_id, first_name, last_name, phone_number, gender, date_of_birth, passport_number, identification_number, address FROM TicketOwnerSnapshots
WHERE ticket_id = $1
LIMIT 1
`

func (q *Queries) GetTicketOwnerSnapshot(ctx context.Context, ticketID int64) (Ticketownersnapshot, error) {
	row := q.db.QueryRow(ctx, getTicketOwnerSnapshot, ticketID)
	var i Ticketownersnapshot
	err := row.Scan(
		&i.TicketID,
		&i.FirstName,
		&i.LastName,
		&i.PhoneNumber,
		&i.Gender,
		&i.DateOfBirth,
		&i.PassportNumber,
		&i.IdentificationNumber,
		&i.Address,
	)
	return i, err
}

const listTicketOwnerSnapshots = `-- name: ListTicketOwnerSnapshots :many
SELECT ticket_id, first_name, last_name, phone_number, gender, date_of_birth, passport_number, identification_number, address FROM TicketOwnerSnapshots
ORDER BY ticket_id
LIMIT $1
OFFSET $2
`

type ListTicketOwnerSnapshotsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListTicketOwnerSnapshots(ctx context.Context, arg ListTicketOwnerSnapshotsParams) ([]Ticketownersnapshot, error) {
	rows, err := q.db.Query(ctx, listTicketOwnerSnapshots, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Ticketownersnapshot{}
	for rows.Next() {
		var i Ticketownersnapshot
		if err := rows.Scan(
			&i.TicketID,
			&i.FirstName,
			&i.LastName,
			&i.PhoneNumber,
			&i.Gender,
			&i.DateOfBirth,
			&i.PassportNumber,
			&i.IdentificationNumber,
			&i.Address,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
