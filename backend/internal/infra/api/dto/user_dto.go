package dto

import "time"

type UserGetResponse struct {
	UserID    int64  `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type UserCreateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserUpdateRequest struct {
	FirstName            string    `json:"firstName"`
	LastName             string    `json:"lastName"`
	PhoneNumber          string    `json:"phoneNumber"`
	DateOfBirth          time.Time `json:"dateOfBirth"`
	Gender               string    `json:"gender"`
	Address              string    `json:"address"`
	PassportNumber       string    `json:"passportNumber"`
	IdentificationNumber string    `json:"identificationNumber"`
	BookingHistory       []string  `json:"bookingHistory"`
	LoyaltyPoints        int64     `json:"loyaltyPoints"`
}

type UserGetResponseByToken struct {
	UID                  int64     `json:"uid"`
	FirstName            string    `json:"firstName"`
	LastName             string    `json:"lastName"`
	Email                string    `json:"email"`
	PhoneNumber          string    `json:"phoneNumber"`
	DateOfBirth          time.Time `json:"dateOfBirth"`
	Gender               string    `json:"gender"`
	Address              string    `json:"address"`
	PassportNumber       string    `json:"passportNumber"`
	IdentificationNumber string    `json:"identificationNumber"`
	LoyaltyPoints        int64     `json:"LoyaltyPoints"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"UpdatedAt"`
}
