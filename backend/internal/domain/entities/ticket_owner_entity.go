package entities

import "time"

type GenderType string

const (
	GenderMale   GenderType = "Male"
	GenderFemale GenderType = "Female"
	GenderOther  GenderType = "Other"
)

type TicketOwner struct {
	TicketID             int64      `json:"ticket_id"`
	FirstName            string     `json:"first_name"`
	LastName             string     `json:"last_name"`
	PhoneNumber          string     `json:"phone_number"`
	Gender               GenderType `json:"gender"`
	DateOfBirth          time.Time  `json:"date_of_birth"`
	PassportNumber       string     `json:"passport_number"`
	IdentificationNumber string     `json:"identification_number"`
	Address              string     `json:"address"`
}
