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
