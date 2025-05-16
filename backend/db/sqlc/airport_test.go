package db

import (
	"context"
	"testing"
	"time"

	"github.com/spaghetti-lover/qairlines/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAirport(t *testing.T) Airport {
	arg := CreateAirportParams{
		AirportCode: utils.RandomStringNum(),
		City:        utils.RandomString(6),
		Name:        utils.RandomName(),
	}

	airport, err := testStore.CreateAirport(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, airport)
	require.NotEmpty(t, airport.AirportID)
	require.NotEmpty(t, airport.AirportCode)
	require.NotEmpty(t, airport.Name)
	require.NotEmpty(t, airport.City)
	require.NotEmpty(t, airport.CreatedAt)

	return airport

}

func TestCreateAirport(t *testing.T) {
	createRandomAirport(t)
}

func TestGetAirport(t *testing.T) {
	airport1 := createRandomAirport(t)
	airport2, err := testStore.GetAirport(context.Background(), airport1.AirportCode)
	require.NoError(t, err)
	require.NotEmpty(t, airport2)

	require.Equal(t, airport1.AirportCode, airport2.AirportCode)
	require.Equal(t, airport1.AirportID, airport2.AirportID)
	require.Equal(t, airport1.City, airport2.City)
	require.Equal(t, airport1.Name, airport2.Name)
	require.WithinDuration(t, airport1.CreatedAt, airport2.CreatedAt, time.Second)
}

func TestDeleteAirport(t *testing.T) {
	airport1 := createRandomAirport(t)
	err := testStore.DeleteAirport(context.Background(), airport1.AirportCode)
	require.NoError(t, err)

	airport2, err := testStore.GetAirport(context.Background(), airport1.AirportCode)
	require.Error(t, err)
	require.EqualError(t, err, "no rows in result set")
	require.Empty(t, airport2)
}

func TestListAirport(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAirport(t)
	}

	arg := ListAirportsParams{
		Limit:  5,
		Offset: 5,
	}

	airports, err := testStore.ListAirports(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, airports, 5)

	for _, airport := range airports {
		require.NotEmpty(t, airport)
	}
}
