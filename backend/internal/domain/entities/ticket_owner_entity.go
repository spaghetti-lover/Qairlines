package entities

type TicketOwner struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
}
