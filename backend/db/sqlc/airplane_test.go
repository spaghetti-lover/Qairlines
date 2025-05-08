package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/spaghetti-lover/qairlines/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAirplane(t *testing.T) Airplane {
	airplane_model := createRandomAirplaneModel(t)
	arg := CreateAirplaneParams{
		AirplaneModelID:    airplane_model.AirplaneModelID,
		RegistrationNumber: utils.RandomStringNum(),
	}

	airplane, err := testQueries.CreateAirplane(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, airplane)
	require.NotEmpty(t, airplane.AirplaneModelID)
	require.NotEmpty(t, airplane.RegistrationNumber)
	require.True(t, airplane.Active.Bool)

	return airplane

}

func TestCreateAirplane(t *testing.T) {
	createRandomAirplane(t)
}

func TestGetAirplane(t *testing.T) {
	airplane1 := createRandomAirplane(t)
	airplane2, err := testQueries.GetAirplane(context.Background(), airplane1.RegistrationNumber)
	require.NoError(t, err)
	require.NotEmpty(t, airplane2)

	require.Equal(t, airplane1.AirplaneID, airplane2.AirplaneID)
	require.Equal(t, airplane1.RegistrationNumber, airplane2.RegistrationNumber)
	require.Equal(t, airplane1.AirplaneModelID, airplane2.AirplaneModelID)
	require.Equal(t, airplane1.Active, airplane1.Active)
}

func TestDeleteAirplane(t *testing.T) {
	airplane1 := createRandomAirplane(t)
	err := testQueries.DeleteAirplane(context.Background(), airplane1.RegistrationNumber)
	require.NoError(t, err)

	airplane2, err := testQueries.GetAirplane(context.Background(), airplane1.RegistrationNumber)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, airplane2)
}

func TestListAirplanes(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAirplane(t)
	}

	arg := ListAirplanesParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAirplanes(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
