package dto

type GetTicketByFlightIDResponse struct {
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

type GetTicketResponse struct {
	TicketID    int64               `json:"ticketId"`
	Status      string              `json:"status"`
	SeatCode    string              `json:"seatCode"`
	FlightClass string              `json:"flightClass"`
	Price       int                 `json:"price"`
	OwnerData   TicketOwnerResponse `json:"ownerData"`
	BookingID   int64               `json:"bookingId"`
	FlightID    int64               `json:"flightId"`
	CreatedAt   string              `json:"createdAt"`
	UpdatedAt   string              `json:"updatedAt"`
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

type CancelTicketResponse struct {
	TicketID    int64               `json:"ticketId"`
	Status      string              `json:"status"`
	SeatCode    string              `json:"seatCode"`
	FlightClass string              `json:"flightClass"`
	Price       int                 `json:"price"`
	OwnerData   TicketOwnerResponse `json:"ownerData"`
	BookingID   int64               `json:"bookingId"`
	FlightID    int64               `json:"flightId"`
	UpdatedAt   string              `json:"updatedAt"`
}

type UpdateSeatRequest struct {
	TicketID string  `json:"ticketId"`
	SeatCode string `json:"seatCode"`
}

type UpdateSeatResponse struct {
	TicketID int64  `json:"ticketId"`
	SeatCode string `json:"seatCode"`
	Status   string `json:"status"`
}
