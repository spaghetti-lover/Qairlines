package entities

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
