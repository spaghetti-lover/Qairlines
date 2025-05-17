package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/spaghetti-lover/qairlines/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestBookingTx(t *testing.T) {
	flight := createRandomFlightWithSeats(t)
	arg := BookingTxParams{
		BookerEmail:      utils.RandomEmail(),
		NumberOfAdults:   1,
		NumberOfChildren: 0,
		FlightClass:      "Economy",
		Cancelled:        pgtype.Bool{Bool: false, Valid: true},
		FlightID:         flight.FlightID,
	}

	ctx := context.Background()

	// Run n cocurrent booking transaction
	n := 5
	errs := make(chan error)
	results := make(chan BookingTxResult)
	occupiedBefore, _ := testStore.CountOccupiedSeats(ctx, CountOccupiedSeatsParams{
		FlightID:    arg.FlightID,
		FlightClass: arg.FlightClass,
	})
	for i := 0; i < n; i++ {
		go func() {
			passengers := []PassengerParams{
				{
					CitizenID:      "123456789",
					PassportNumber: pgtype.Text{String: "P123456", Valid: true},
					Gender:         "Male",
					PhoneNumber:    "0123456789",
					FirstName:      "John",
					LastName:       "Doe",
					Nationality:    "VN",
					DateOfBirth:    pgtype.Date{Time: time.Now().AddDate(-25, 0, 0), Valid: true},
					SeatRow:        int32(utils.RandomInt(1, 50)),
					SeatCol:        string(rune('A' + utils.RandomInt(0, 50))),
				},
			}

			payment := PaymentParams{
				Amount:        utils.RandomPrice(),
				Currency:      pgtype.Text{String: "VND", Valid: true},
				PaymentMethod: pgtype.Text{String: "credit_card", Valid: true},
				Status:        pgtype.Text{String: "paid", Valid: true},
			}
			result, err := testStore.BookingTx(ctx, arg, passengers, payment)
			errs <- err
			results <- result
		}()
	}
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		// Check booking
		result := <-results
		require.NotEmpty(t, result)
		require.Equal(t, arg.BookerEmail, result.BookerEmail)
		require.Equal(t, arg.NumberOfAdults, result.NumberOfAdults)
		require.Equal(t, arg.FlightClass, result.FlightClass)
		require.Equal(t, arg.NumberOfChildren, result.NumberOfChildren)
		require.Equal(t, arg.FlightClass, result.FlightClass)
		require.Equal(t, arg.Cancelled, result.Cancelled)
		require.Equal(t, arg.FlightID, result.FlightID)

		// Check flight_seat
		// 		- Them truong hop tao flight nhung ma arrival phai muon hon start time
	}
	occupiedAfter, _ := testStore.CountOccupiedSeats(ctx, CountOccupiedSeatsParams{
		FlightID:    arg.FlightID,
		FlightClass: arg.FlightClass,
	})
	require.Equal(t, int64(n), occupiedAfter-occupiedBefore)
}

func createRandomFlightWithSeats(t *testing.T) Flight {
	ctx := context.Background()

	departureAirport := createRandomAirport(t)
	destinationAirport := createRandomAirport(t)

	airplane := createRandomAirplane(t)
	flightParams := CreateFlightParams{
		FlightNumber:           utils.RandomStringNum(),
		RegistrationNumber:     airplane.RegistrationNumber,
		EstimatedDepartureTime: utils.RandomTime(),
		EstimatedArrivalTime:   utils.RandomTime(),
		DepartureAirportID:     departureAirport.AirportID,
		DestinationAirportID:   destinationAirport.AirportID,
		FlightPrice:            utils.RandomPrice(),
		Status:                 "Scheduled",
	}

	params := CreateFlightWithSeatsParams{
		FlightParams: flightParams,
		EconomyParams: CreateFlightClassParams{
			MaxRow:     100,
			MaxCol:     200,
			Multiplier: 1.0,
		},
		BusinessParams: CreateFlightClassParams{
			MaxRow:     100,
			MaxCol:     200,
			Multiplier: 1.5,
		},
		FirstParams: CreateFlightClassParams{
			MaxRow:     100,
			MaxCol:     200,
			Multiplier: 2.0,
		},
	}

	flight, err := testStore.CreateFlightWithSeats(ctx, params)
	require.NoError(t, err)
	require.NotEmpty(t, flight)
	require.Equal(t, params.FlightParams.FlightNumber, flight.FlightNumber)

	seats, err := testStore.ListFlightSeatsByFlightID(ctx, flight.FlightID)
	require.NoError(t, err)
	require.Len(t, seats, 3)

	classExists := map[string]bool{"Economy": false, "Business": false, "First": false}
	for _, seat := range seats {
		classExists[string(seat.FlightClass)] = true
		require.Equal(t, flight.FlightID, seat.FlightID)
	}
	require.True(t, classExists["Economy"])
	require.True(t, classExists["Business"])
	require.True(t, classExists["First"])
	return flight

}

// func TestCreateFlightWithSeats(t *testing.T) {
// 	createRandomFlightWithSeats(t)
// }
