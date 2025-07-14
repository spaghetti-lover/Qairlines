package dto

import "github.com/spaghetti-lover/qairlines/internal/domain/entities"

type CreateCustomerRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type CreateCustomerResponse struct {
	Message string `json:"message"`
	User    struct {
		ID        int64  `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
	} `json:"user"`
}

type DateOfBirth struct {
	Seconds int64 `json:"seconds"`
}

type CustomerUpdateRequest struct {
	FirstName            string      `json:"firstName"`
	LastName             string      `json:"lastName"`
	PhoneNumber          string      `json:"phoneNumber"`
	Gender               string      `json:"gender"`
	Address              string      `json:"address"`
	PassportNumber       string      `json:"passportNumber"`
	IdentificationNumber string      `json:"identificationNumber"`
	DateOfBirth          DateOfBirth `json:"dateOfBirth"`
	UID                  string      `json:"uid"`
}

type CustomerUpdateResponse struct {
	UID                  int64               `json:"uid"`
	FirstName            string              `json:"firstName"`
	LastName             string              `json:"lastName"`
	PhoneNumber          string              `json:"phoneNumber"`
	Gender               entities.GenderType `json:"gender"`
	DateOfBirth          DateOfBirth         `json:"dateOfBirth"`
	Address              string              `json:"address"`
	Passport             string              `json:"passport"`
	IdentificationNumber string              `json:"identificationNumber"`
}

type CustomerResponse struct {
	UID                  string      `json:"uid"`
	FirstName            string      `json:"firstName"`
	LastName             string      `json:"lastName"`
	Email                string      `json:"email"`
	DateOfBirth          TimeSeconds `json:"dateOfBirth"`
	Gender               string      `json:"gender"`
	LoyaltyPoints        int32       `json:"loyaltyPoints"`
	CreatedAt            TimeSeconds `json:"createdAt"`
	Address              string      `json:"address"`
	PassportNumber       string      `json:"passportNumber"`
	IdentificationNumber string      `json:"identificationNumber"`
}

type CustomerDetailsResponse struct {
	UID                  string       `json:"uid"`
	Role                 string       `json:"role"`
	PhoneNumber          string       `json:"phoneNumber"`
	DateOfBirth          string       `json:"dateOfBirth"`
	FirstName            string       `json:"firstName"`
	LastName             string       `json:"lastName"`
	Gender               string       `json:"gender"`
	Email                string       `json:"email"`
	IdentificationNumber *string      `json:"identificationNumber"`
	PassportNumber       string       `json:"passportNumber"`
	Address              string       `json:"address"`
	LoyaltyPoints        int          `json:"loyaltyPoints"`
	BookingHistory       []string     `json:"bookingHistory"`
	CreatedAt            TimeWithNano `json:"createdAt"`
	UpdatedAt            TimeWithNano `json:"updatedAt"`
}

type TimeWithNano struct {
	Seconds     int64 `json:"seconds"`
	Nanoseconds int64 `json:"nanoseconds"`
}

type ListCustomersParams struct {
	Limit int `json:"limit" binding:"required,min=1,max=100" default:"10"`
	Page  int `json:"page" binding:"required,min=1" default:"1"`
}
