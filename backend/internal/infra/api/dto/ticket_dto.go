package dto

type GetTicketResponse struct {
	TicketID    int64               `json:"ticketId"`
	SeatID      int64               `json:"seatId"`
	FlightClass string              `json:"flightClass"`
	Price       int                 `json:"price"`
	Status      string              `json:"status"`
	BookingID   int64               `json:"bookingId"`
	FlightID    int64               `json:"flightId"`
	CreatedAt   string              `json:"createdAt"`
	UpdatedAt   string              `json:"updatedAt"`
	Seat        SeatResponse        `json:"seat"`
	Owner       TicketOwnerResponse `json:"owner"`
}

type SeatResponse struct {
	SeatID      int64  `json:"seatId"`
	SeatCode    string `json:"seatCode"`
	IsAvailable bool   `json:"isAvailable"`
	Class       string `json:"class"`
}

type TicketOwnerResponse struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
	Gender      string `json:"gender"`
}
