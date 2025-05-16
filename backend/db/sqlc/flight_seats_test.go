package db

import (
	"context"
	"math/big"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomFlightSeat(t *testing.T) FlightSeat {
	flight := createRandomFlight(t)
	arg := CreateFlightSeatParams{
		FlightID:        flight.FlightID,
		FlightClass:     "Economy",
		ClassMultiplier: pgtype.Numeric{Int: big.NewInt(100), Exp: -2, Valid: true}, // 1.00
		ChildMultiplier: pgtype.Numeric{Int: big.NewInt(50), Exp: -2, Valid: true},  // 0.50
		MaxRowSeat:      200,
		MaxColSeat:      100,
	}

	flightSeat, err := testStore.CreateFlightSeat(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, flightSeat)
	require.Equal(t, arg.FlightClass, flightSeat.FlightClass)

	return flightSeat
}

func TestCreateFlightSeat(t *testing.T) {
	createRandomFlightSeat(t)
}

func TestGetFlightSeat(t *testing.T) {
	fs1 := createRandomFlightSeat(t)

	fs2, err := testStore.GetFlightSeat(context.Background(), GetFlightSeatParams{
		FlightID:    fs1.FlightID,
		FlightClass: fs1.FlightClass,
	})
	require.NoError(t, err)
	require.NotEmpty(t, fs2)
	require.Equal(t, fs1.FlightSeatsID, fs2.FlightSeatsID)
	require.Equal(t, fs1.FlightClass, fs2.FlightClass)
	require.Equal(t, fs1.MaxRowSeat, fs2.MaxRowSeat)
	require.Equal(t, fs1.MaxColSeat, fs2.MaxColSeat)
}

func TestDeleteFlightSeat(t *testing.T) {
	fs := createRandomFlightSeat(t)

	err := testStore.DeleteFlightSeat(context.Background(), fs.FlightID)
	require.NoError(t, err)

	_, err = testStore.GetFlightSeat(context.Background(), GetFlightSeatParams{
		FlightID:    fs.FlightID,
		FlightClass: fs.FlightClass,
	})
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

	seats, err := testStore.ListFlightSeats(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, seats, 3)

	for _, seat := range seats {
		require.NotEmpty(t, seat)
	}
}
