package entities

import "time"

type UserRole string

const (
	RoleAdmin  UserRole = "admin"
	RoleClient UserRole = "customer"
)

type User struct {
	UserID               int64
	FirstName            string
	LastName             string
	PhoneNumber          string
	DateOfBirth          time.Time
	Gender               string
	Address              string
	PassportNumber       string
	IdentificationNumber string
	HashedPassword       string
	Role                 UserRole
	Email                string
	LoyaltyPoints        int64
	UpdatedAt            time.Time
	CreatedAt            time.Time
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
