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
	UserID         int64
	Username       string
	HashedPassword string
	Role           UserRole
}

type ListUsersParams struct {
	Limit  int32
	Offset int32
}

type CreateUserParams struct {
	Username string
	Password string
	Role     UserRole
}

func NewUser(username string, password string, role string) (*User, error) {
	if username == "" {
		return nil, errors.New("username is required")
	}
	if password == "" {
		return nil, errors.New("password is required")
	}
	if role != string(RoleAdmin) && role != string(RoleClient) {
		return nil, errors.New("role must be either 'admin' or 'client'")
	}
	return &User{
		Username:       username,
		HashedPassword: password,
		Role:           UserRole(role),
	}, nil
}
