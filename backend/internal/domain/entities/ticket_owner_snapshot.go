package entities

// ticket_id VARCHAR(30) PRIMARY KEY REFERENCES Tickets(ticket_id) ON DELETE CASCADE,
// first_name VARCHAR(100),
// last_name VARCHAR(100),
// phone_number VARCHAR(20),
// gender gender_type NOT NULL DEFAULT 'other'
type GenderType string

const (
	GenderMale   GenderType = "male"
	GenderFemale GenderType = "female"
	GenderOther  GenderType = "other"
)

type TicketOwnerSnapshot struct {
	TicketID    string     `json:"ticket_id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	PhoneNumber string     `json:"phone_number"`
	Gender      GenderType `json:"gender"`
}
