package entities

import "time"

type UserRole string

const (
	RoleAdmin    UserRole = "admin"
	RoleCustomer UserRole = "customer"
)

type User struct {
	UserID    int64     `json:"user_id"`
	Email     string    `json:"email"`
	HashedPwd string    `json:"hashed_password"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Role      UserRole  `json:"role"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

type UpdateUserParams struct {
	UserID    int64
	FirstName string
	LastName  string
}
