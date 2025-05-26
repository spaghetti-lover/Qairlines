package dto

type TicketOwnerData struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
	Gender      string `json:"gender"`
}

type TicketResponse struct {
	TicketID    string          `json:"ticketId"`
	Status      string          `json:"status"`
	SeatCode    string          `json:"seatCode"`
	FlightClass string          `json:"flightClass"`
	Price       int32           `json:"price"`
	OwnerData   TicketOwnerData `json:"ownerData"`
	BookingID   string          `json:"bookingId"`
	FlightID    string          `json:"flightId"`
	CreatedAt   string          `json:"createdAt"`
	UpdatedAt   string          `json:"updatedAt"`
}

type GetTicketsResponse struct {
	Message string           `json:"message"`
	Data    []TicketResponse `json:"data"`
}
