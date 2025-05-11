package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/spaghetti-lover/qairlines/utils"
	"github.com/stretchr/testify/require"
)

// Hàm tạo booking ngẫu nhiên phục vụ cho các test khác
func createRandomBooking(t *testing.T) Booking {
	flight := createRandomFlight(t)
	arg := CreateBookingParams{
		BookerEmail:      utils.RandomEmail(),
		NumberOfAdults:   2,
		NumberOfChildren: 1,
		FlightClass:      FlightClassTypeEconomy,
		Cancelled:        pgtype.Bool{Bool: false, Valid: true},
		FlightID:         flight.FlightID,
	}

	booking, err := testStore.CreateBooking(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, booking)

	require.Equal(t, arg.BookerEmail, booking.BookerEmail)
	require.Equal(t, arg.NumberOfAdults, booking.NumberOfAdults)
	require.Equal(t, arg.NumberOfChildren, booking.NumberOfChildren)
	require.Equal(t, arg.FlightClass, booking.FlightClass)
	require.Equal(t, arg.Cancelled.Bool, booking.Cancelled.Bool)
	require.Equal(t, arg.FlightID, booking.FlightID)
	require.WithinDuration(t, time.Now(), booking.BookingDate.Time, time.Second*5)

	return booking
}

func TestCreateBooking(t *testing.T) {
	createRandomBooking(t)
}

func TestGetBooking(t *testing.T) {
	booking1 := createRandomBooking(t)
	booking2, err := testStore.GetBooking(context.Background(), booking1.BookingID)
	require.NoError(t, err)
	require.NotEmpty(t, booking2)

	require.Equal(t, booking1.BookingID, booking2.BookingID)
	require.Equal(t, booking1.BookerEmail, booking2.BookerEmail)
	require.Equal(t, booking1.FlightID, booking2.FlightID)
}

func TestDeleteBooking(t *testing.T) {
	booking := createRandomBooking(t)

	err := testStore.DeleteBookings(context.Background(), booking.BookingID)
	require.NoError(t, err)

	_, err = testStore.GetBooking(context.Background(), booking.BookingID)
	require.Error(t, err)
	require.EqualError(t, err, "no rows in result set")
}

func TestListBookings(t *testing.T) {
	// Tạo 10 booking
	for i := 0; i < 10; i++ {
		createRandomBooking(t)
	}

	arg := ListBookingsParams{
		Limit:  5,
		Offset: 5,
	}

	bookings, err := testStore.ListBookings(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, bookings, 5)

	for _, b := range bookings {
		require.NotEmpty(t, b)
	}
}
