package db

import (
	"context"
	"testing"
	"time"

	"github.com/spaghetti-lover/qairlines/pkg/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAirplaneModel(t *testing.T) AirplaneModel {
	arg := CreateAirplaneModelParams{
		Name:         utils.RandomName(),
		Manufacturer: utils.RandomName(),
		TotalSeats:   int64(utils.RandomInt(0, 800)),
	}

	airplane_model, err := testStore.CreateAirplaneModel(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, airplane_model)
	require.NotEmpty(t, airplane_model.Name)
	require.NotEmpty(t, airplane_model.Manufacturer)
	require.GreaterOrEqual(t, airplane_model.TotalSeats, int64(0))

	return airplane_model

}

func TestCreateAirplaneModel(t *testing.T) {
	createRandomAirplaneModel(t)
}

func TestGetAirplaneModel(t *testing.T) {
	airplane_model1 := createRandomAirplaneModel(t)
	airplane_model2, err := testStore.GetAirplaneModel(context.Background(), airplane_model1.AirplaneModelID)
	require.NoError(t, err)
	require.NotEmpty(t, airplane_model2)

	require.Equal(t, airplane_model1.AirplaneModelID, airplane_model2.AirplaneModelID)
	require.Equal(t, airplane_model1.Name, airplane_model2.Name)
	require.Equal(t, airplane_model1.Manufacturer, airplane_model2.Manufacturer)
	require.Equal(t, airplane_model1.TotalSeats, airplane_model2.TotalSeats)
	require.WithinDuration(t, airplane_model1.CreatedAt, airplane_model1.CreatedAt, time.Second)
}

func TestDeleteAirplaneModel(t *testing.T) {
	airplane_model1 := createRandomAirplaneModel(t)
	err := testStore.DeleteAirplaneModel(context.Background(), airplane_model1.AirplaneModelID)
	require.NoError(t, err)

	airplane_model2, err := testStore.GetAirplaneModel(context.Background(), airplane_model1.AirplaneModelID)
	require.Error(t, err)
	require.EqualError(t, err, "no rows in result set")
	require.Empty(t, airplane_model2)
}

func TestListAirplaneModels(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAirplaneModel(t)
	}

	arg := ListAirplaneModelsParams{
		Limit:  5,
		Offset: 5,
	}

	airplaneModels, err := testStore.ListAirplaneModels(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, airplaneModels, 5)

	for _, airplaneModel := range airplaneModels {
		require.NotEmpty(t, airplaneModel)
	}
}
