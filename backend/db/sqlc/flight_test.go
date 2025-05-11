package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/spaghetti-lover/qairlines/utils"
	"github.com/stretchr/testify/require"
)

func createRandomFlight(t *testing.T) Flight {
	airplane := createRandomAirplane(t)
	airport1 := createRandomAirport(t)
	airport2 := createRandomAirport(t)
	arg := CreateFlightParams{
		FlightNumber:           utils.RandomStringNum(),
		RegistrationNumber:     airplane.RegistrationNumber,
		EstimatedDepartureTime: utils.RandomTime(),
		ActualDepartureTime:    utils.RandomTime(),
		EstimatedArrivalTime:   utils.RandomTime(),
		ActualArrivalTime:      utils.RandomTime(),
		DepartureAirportID:     airport1.AirportID,
		DestinationAirportID:   airport2.AirportID,
		FlightPrice:            utils.RandomPrice(),
		Status:                 FlightStatusScheduled,
	}

	flight, err := testStore.CreateFlight(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, flight)
	require.NotEmpty(t, flight.FlightID)
	require.NotEmpty(t, flight.DepartureAirportID)
	require.NotEmpty(t, flight.DestinationAirportID)
	require.NotEmpty(t, flight.FlightPrice)
	require.NotEmpty(t, flight.Status)

	return flight

}

func TestCreateFlight(t *testing.T) {
	createRandomFlight(t)
}

func TestGetFlight(t *testing.T) {
	flight1 := createRandomFlight(t)

	flight2, err := testStore.GetFlight(context.Background(), flight1.FlightID)
	require.NoError(t, err)
	require.NotEmpty(t, flight2)

	require.Equal(t, flight1.FlightID, flight2.FlightID)
	require.Equal(t, flight1.FlightNumber, flight2.FlightNumber)
	require.Equal(t, flight1.RegistrationNumber, flight2.RegistrationNumber)

	require.True(t, flight2.EstimatedDepartureTime.Valid)
	require.True(t, flight2.ActualDepartureTime.Valid)
	require.True(t, flight2.EstimatedArrivalTime.Valid)
	require.True(t, flight2.ActualArrivalTime.Valid)

	require.WithinDuration(t, flight1.EstimatedDepartureTime.Time, flight2.EstimatedDepartureTime.Time, time.Second)
	require.WithinDuration(t, flight1.ActualDepartureTime.Time, flight2.ActualDepartureTime.Time, time.Second)
	require.WithinDuration(t, flight1.EstimatedArrivalTime.Time, flight2.EstimatedArrivalTime.Time, time.Second)
	require.WithinDuration(t, flight1.ActualArrivalTime.Time, flight2.ActualArrivalTime.Time, time.Second)

	require.Equal(t, flight1.DepartureAirportID, flight2.DepartureAirportID)
	require.Equal(t, flight1.DestinationAirportID, flight2.DestinationAirportID)

	require.True(t, flight2.FlightPrice.Valid)
	require.Equal(t, flight1.FlightPrice.Int.String(), flight2.FlightPrice.Int.String())
	require.Equal(t, flight1.FlightPrice.Exp, flight2.FlightPrice.Exp)

	require.Equal(t, flight1.Status, flight2.Status)
}

func TestDeleteFlight(t *testing.T) {
	flight1 := createRandomFlight(t)

	err := testStore.DeleteFlight(context.Background(), flight1.FlightNumber)
	require.NoError(t, err)

	flight2, err := testStore.GetFlight(context.Background(), flight1.FlightID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, flight2)
}
func TestDeleteNonExistentFlight(t *testing.T) {
	nonExistentFlightCode := "FLIGHT-404-NOTFOUND"
	err := testStore.DeleteFlight(context.Background(), nonExistentFlightCode)
	require.NoError(t, err)
}

func TestListFlights(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomFlight(t)
	}

	arg := ListFlightsParams{
		Limit:  5,
		Offset: 5,
	}

	flights, err := testStore.ListFlights(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, flights, 5)

	for _, Flight := range flights {
		require.NotEmpty(t, Flight)
	}
}
