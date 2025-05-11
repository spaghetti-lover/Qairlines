package db

import (
	"context"
	"fmt"
	"math/big"

	"github.com/jackc/pgx/v5/pgtype"
)

type CreateFlightClassParams struct {
	MaxRow     int64   `json:"max_row"`
	MaxCol     int64   `json:"max_col"`
	Multiplier float64 `json:"multiplier"`
}

type CreateFlightWithSeatsParams struct {
	FlightParams   CreateFlightParams      `json:"flight"`
	EconomyParams  CreateFlightClassParams `json:"economy"`
	BusinessParams CreateFlightClassParams `json:"business"`
	FirstParams    CreateFlightClassParams `json:"first"`
}

func (store *SQLStore) CreateFlightWithSeats(ctx context.Context, params CreateFlightWithSeatsParams) (Flight, error) {
	var flight Flight

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		flight, err = q.CreateFlight(ctx, params.FlightParams)
		if err != nil {
			return fmt.Errorf("create flight failed: %w", err)
		}

		seatsConfig := map[FlightClassType]CreateFlightClassParams{
			"Economy":  params.EconomyParams,
			"Business": params.BusinessParams,
			"First":    params.FirstParams,
		}

		for class, config := range seatsConfig {
			_, err = q.CreateFlightSeat(ctx, CreateFlightSeatParams{
				FlightID:           flight.FlightID,
				RegistrationNumber: flight.RegistrationNumber,
				FlightClass:        class,
				ClassMultiplier:    pgtype.Numeric{Int: big.NewInt(int64(config.Multiplier * 100)), Exp: -2, Valid: true},
				ChildMultiplier:    pgtype.Numeric{Int: big.NewInt(75), Exp: -2, Valid: true},
				MaxRowSeat:         config.MaxRow,
				MaxColSeat:         config.MaxCol,
			})
			if err != nil {
				return fmt.Errorf("create flight_seats for class %s failed: %w", class, err)
			}
		}

		return nil
	})

	return flight, err
}
