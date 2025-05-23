// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: airport.sql

package db

import (
	"context"
)

const createAirport = `-- name: CreateAirport :one
INSERT INTO airport (
  airport_code,
  city,
  name
) VALUES (
  $1, $2, $3
) RETURNING airport_id, airport_code, city, name, created_at
`

type CreateAirportParams struct {
	AirportCode string `json:"airport_code"`
	City        string `json:"city"`
	Name        string `json:"name"`
}

func (q *Queries) CreateAirport(ctx context.Context, arg CreateAirportParams) (Airport, error) {
	row := q.db.QueryRow(ctx, createAirport, arg.AirportCode, arg.City, arg.Name)
	var i Airport
	err := row.Scan(
		&i.AirportID,
		&i.AirportCode,
		&i.City,
		&i.Name,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAirport = `-- name: DeleteAirport :exec
DELETE FROM airport
WHERE airport_code = $1
`

func (q *Queries) DeleteAirport(ctx context.Context, airportCode string) error {
	_, err := q.db.Exec(ctx, deleteAirport, airportCode)
	return err
}

const getAirport = `-- name: GetAirport :one
SELECT airport_id, airport_code, city, name, created_at FROM airport
WHERE airport_code = $1 LIMIT 1
`

func (q *Queries) GetAirport(ctx context.Context, airportCode string) (Airport, error) {
	row := q.db.QueryRow(ctx, getAirport, airportCode)
	var i Airport
	err := row.Scan(
		&i.AirportID,
		&i.AirportCode,
		&i.City,
		&i.Name,
		&i.CreatedAt,
	)
	return i, err
}

const listAirports = `-- name: ListAirports :many
SELECT airport_id, airport_code, city, name, created_at FROM airport
ORDER BY airport_code
LIMIT $1
OFFSET $2
`

type ListAirportsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAirports(ctx context.Context, arg ListAirportsParams) ([]Airport, error) {
	rows, err := q.db.Query(ctx, listAirports, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Airport{}
	for rows.Next() {
		var i Airport
		if err := rows.Scan(
			&i.AirportID,
			&i.AirportCode,
			&i.City,
			&i.Name,
			&i.CreatedAt,
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
