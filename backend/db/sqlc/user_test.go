package db

import (
	"context"
	"testing"

	"github.com/spaghetti-lover/qairlines/pkg/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		FirstName:      utils.RandomName(),
		LastName:       utils.RandomName(),
		HashedPassword: utils.RandomName(),
		Email:          utils.RandomEmail(),
	}

	user, err := testStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.NotEmpty(t, user.FirstName)
	require.NotEmpty(t, user.LastName)
	require.NotEmpty(t, user.HashedPassword)
	require.NotEmpty(t, user.Role)

	return user

}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testStore.GetUser(context.Background(), user1.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.UserID, user2.UserID)
	require.Equal(t, user1.FirstName, user2.FirstName)
	require.Equal(t, user1.LastName, user2.LastName)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
}

func TestDeleteUser(t *testing.T) {
	User1 := createRandomUser(t)
	err := testStore.DeleteUser(context.Background(), User1.UserID)
	require.NoError(t, err)

	User2, err := testStore.GetUser(context.Background(), User1.UserID)
	require.Error(t, err)
	require.EqualError(t, err, "no rows in result set")
	require.Empty(t, User2)
}

func TestListUser(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}

	Users, err := testStore.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, Users, 5)

	for _, User := range Users {
		require.NotEmpty(t, User)
	}
}

func TestGetAllUser(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}
	users, err := testStore.GetAllUser(context.Background())
	require.NotEmpty(t, users)
	require.NoError(t, err)
	for _, user := range users {
		require.NotEmpty(t, user)
	}
}
