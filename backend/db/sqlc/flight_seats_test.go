package db

import (
	"context"
	"math/big"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomFlightSeat(t *testing.T) FlightSeat {
	airplane := createRandomAirplane(t)

	arg := CreateFlightSeatParams{
		RegistrationNumber: airplane.RegistrationNumber,
		FlightClass:        "Economy",
		ClassMultiplier:    pgtype.Numeric{Int: big.NewInt(100), Exp: -2, Valid: true}, // 1.00
		ChildMultiplier:    pgtype.Numeric{Int: big.NewInt(50), Exp: -2, Valid: true},  // 0.50
		MaxRowSeat:         30,
		MaxColSeat:         6,
	}

	flightSeat, err := testQueries.CreateFlightSeat(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, flightSeat)
	require.Equal(t, arg.RegistrationNumber, flightSeat.RegistrationNumber)
	require.Equal(t, arg.FlightClass, flightSeat.FlightClass)

	return flightSeat
}

func TestCreateFlightSeat(t *testing.T) {
	createRandomFlightSeat(t)
}

func TestGetFlightSeat(t *testing.T) {
	fs1 := createRandomFlightSeat(t)

	fs2, err := testQueries.GetFlightSeat(context.Background(), fs1.RegistrationNumber)
	require.NoError(t, err)
	require.NotEmpty(t, fs2)
	require.Equal(t, fs1.FlightSeatsID, fs2.FlightSeatsID)
	require.Equal(t, fs1.FlightClass, fs2.FlightClass)
	require.Equal(t, fs1.MaxRowSeat, fs2.MaxRowSeat)
	require.Equal(t, fs1.MaxColSeat, fs2.MaxColSeat)
}

func TestDeleteFlightSeat(t *testing.T) {
	fs := createRandomFlightSeat(t)

	err := testQueries.DeleteFlightSeat(context.Background(), fs.RegistrationNumber)
	require.NoError(t, err)

	_, err = testQueries.GetFlightSeat(context.Background(), fs.RegistrationNumber)
	require.Error(t, err)
}

func TestListFlightSeats(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomFlightSeat(t)
	}

	arg := ListFlightSeatsParams{
		Limit:  3,
		Offset: 1,
	}

	seats, err := testQueries.ListFlightSeats(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, seats, 3)

	for _, seat := range seats {
		require.NotEmpty(t, seat)
	}
}
