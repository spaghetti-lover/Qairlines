package entities

import (
	"errors"
)

type UserRole string

const (
	RoleAdmin  UserRole = "admin"
	RoleClient UserRole = "client"
)

type User struct {
	UserID            int64
	FirstName         string
	LastName          string
	HashedPassword    string
	Role              UserRole
	Email             string
	PasswordChangedAt string
	CreatedAt         string
}

type ListUsersParams struct {
	Limit  int32
	Offset int32
}

type CreateUserParams struct {
	FirstName string
	LastName  string
	Password  string
	Email     string
}

func NewUser(firstname string, lastname string, password string, role string) (*User, error) {
	if firstname == "" {
		return nil, errors.New("username is required")
	}
	if lastname == "" {
		return nil, errors.New("lastname is required")
	}
	if password == "" {
		return nil, errors.New("password is required")
	}
	if role != string(RoleAdmin) && role != string(RoleClient) {
		return nil, errors.New("role must be either 'admin' or 'client'")
	}
	return &User{
		FirstName:      firstname,
		LastName:       lastname,
		HashedPassword: password,
		Role:           UserRole(role),
	}, nil
}
