package entities

import "time"

type CustomerGender string

const (
	CustomerGenderMale   CustomerGender = "Male"
	CustomerGenderFemale CustomerGender = "Female"
	CustomerGenderOther  CustomerGender = "Other"
)

type Customer struct {
	UserID               int64          `json:"user_id"`
	PhoneNumber          string         `json:"phone_number"`
	Gender               CustomerGender `json:"gender"`
	DateOfBirth          time.Time      `json:"date_of_birth"`
	PassportNumber       string         `json:"passport_number"`
	IdentificationNumber string         `json:"identification_number"`
	Address              string         `json:"address"`
	LoyaltyPoints        int32          `json:"loyalty_points"`
	User                 User           `json:"user"`
}

type CreateCustomerParams struct {
	UserID               int64          `json:"user_id"`
	PhoneNumber          string         `json:"phone_number"`
	Gender               CustomerGender `json:"gender"`
	DateOfBirth          time.Time      `json:"date_of_birth"`
	PassportNumber       string         `json:"passport_number"`
	IdentificationNumber string         `json:"identification_number"`
	Address              string         `json:"address"`
	LoyaltyPoints        int32          `json:"loyalty_points"`
}
