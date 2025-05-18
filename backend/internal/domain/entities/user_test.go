package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	t.Run("should create an user and return it with error nil", func(t *testing.T) {
		user, err := NewUser("test user", "testuser@testmail.com")

		assert.Nil(t, err)
		assert.NotNil(t, user.ID)
		assert.Equal(t, "test user", user.Name)
		assert.Equal(t, "testuser@testmail.com", user.Email)
	})

	t.Run("should return error for invalid email", func(t *testing.T) {
		_, err := NewUser("test user", "testusertestmail.com")

		assert.Equal(t, "invalid email address", err.Error())
	})
}

func TestValidMailAddress(t *testing.T) {
	t.Run("should return the email and error nil with simple email format", func(t *testing.T) {
		email, err := validMailAddress("testuser@testmail.com")

		assert.Nil(t, err)
		assert.Equal(t, "testuser@testmail.com", email)
	})

	t.Run("should return the email and error nil with RFC 5322 format", func(t *testing.T) {
		email, err := validMailAddress("Test User <testuser@testmail.com>")

		assert.Nil(t, err)
		assert.Equal(t, "testuser@testmail.com", email)
	})

	t.Run("should return error if email is not valid", func(t *testing.T) {
		_, err := validMailAddress("testusertestmail.com")

		assert.NotNil(t, err)
	})
}
