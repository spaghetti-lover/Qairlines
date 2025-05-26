package entities

import "time"

type CreateAdminParams struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type Admin struct {
	UserID    string    `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
