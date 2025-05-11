package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomPassenger(t *testing.T) Passenger {
	booking := createRandomBooking(t)
	dob := pgtype.Date{}
	err := dob.Scan(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC))
	require.NoError(t, err)

	arg := CreatePassengerParams{
		BookingID:      booking.BookingID,
		CitizenID:      "012345678901",
		PassportNumber: pgtype.Text{String: "A1234567", Valid: true},
		Gender:         "M",
		PhoneNumber:    "0901234567",
		FirstName:      "John",
		LastName:       "Doe",
		Nationality:    "Vietnamese",
		DateOfBirth:    dob,
		SeatRow:        12,
		SeatCol:        "C",
	}

	passenger, err := testStore.CreatePassenger(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, passenger)

	require.Equal(t, arg.BookingID, passenger.BookingID)
	require.Equal(t, arg.FirstName, passenger.FirstName)
	require.Equal(t, arg.LastName, passenger.LastName)
	require.Equal(t, arg.SeatRow, passenger.SeatRow)

	return passenger
}

func TestCreatePassenger(t *testing.T) {
	createRandomPassenger(t)
}

func TestGetPassenger(t *testing.T) {
	p1 := createRandomPassenger(t)

	p2, err := testStore.GetPassenger(context.Background(), p1.PassengerID)
	require.NoError(t, err)
	require.NotEmpty(t, p2)

	require.Equal(t, p1.PassengerID, p2.PassengerID)
	require.Equal(t, p1.CitizenID, p2.CitizenID)
}

func TestDeletePassenger(t *testing.T) {
	p := createRandomPassenger(t)

	err := testStore.DeletePassenger(context.Background(), p.PassengerID)
	require.NoError(t, err)

	_, err = testStore.GetPassenger(context.Background(), p.PassengerID)
	require.Error(t, err)
}

func TestListPassengers(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomPassenger(t)
	}

	arg := ListPassengersParams{
		Limit:  3,
		Offset: 1,
	}

	passengers, err := testStore.ListPassengers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, passengers, 3)

	for _, p := range passengers {
		require.NotEmpty(t, p)
	}
}

func TestCountOccupiedSeats(t *testing.T) {
	flight_seat := createRandomFlightSeat(t)
	arg := CountOccupiedSeatsParams{
		FlightID:    flight_seat.FlightID,
		FlightClass: flight_seat.FlightClass,
	}

	count, err := testStore.CountOccupiedSeats(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, count, int64(0))

}
