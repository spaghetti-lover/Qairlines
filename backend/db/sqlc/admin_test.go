package db

import (
	"context"
	"testing"

	"github.com/spaghetti-lover/qairlines/pkg/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAdmin(t *testing.T) Admin {
	arg := CreateAdminParams{
		Username: utils.RandomName(),
		Password: utils.RandomName(),
	}

	admin, err := testStore.CreateAdmin(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, admin)
	require.NotEmpty(t, admin.Username)
	require.NotEmpty(t, admin.Password)

	return admin

}

func TestCreateAdmin(t *testing.T) {
	createRandomAdmin(t)
}

func TestGetAdmin(t *testing.T) {
	admin1 := createRandomAdmin(t)
	admin2, err := testStore.GetAdmin(context.Background(), admin1.AdminID)
	require.NoError(t, err)
	require.NotEmpty(t, admin2)

	require.Equal(t, admin1.AdminID, admin2.AdminID)
	require.Equal(t, admin1.Username, admin2.Username)
	require.Equal(t, admin1.Password, admin2.Password)
}

func TestDeleteAdmin(t *testing.T) {
	admin1 := createRandomAdmin(t)
	err := testStore.DeleteAirplaneModel(context.Background(), admin1.AdminID)
	require.NoError(t, err)

	admin2, err := testStore.GetAirplaneModel(context.Background(), admin1.AdminID)
	require.Error(t, err)
	require.EqualError(t, err, "no rows in result set")
	require.Empty(t, admin2)
}

func TestListAdmin(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAdmin(t)
	}

	arg := ListAdminsParams{
		Limit:  5,
		Offset: 5,
	}

	admins, err := testStore.ListAdmins(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, admins, 5)

	for _, admin := range admins {
		require.NotEmpty(t, admin)
	}
}
