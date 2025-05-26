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
	BookingID   int64           `json:"bookingId"`
	FlightID    int64           `json:"flightId"`
	CreatedAt   string          `json:"createdAt"`
	UpdatedAt   string          `json:"updatedAt"`
}

type GetTicketsResponse struct {
	Message string           `json:"message"`
	Data    []TicketResponse `json:"data"`
}

type OwnerData struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
}

type TicketDetails struct {
	TicketID    string    `json:"ticketId"`
	Status      string    `json:"status"`
	SeatCode    string    `json:"seatCode"`
	FlightClass string    `json:"flightClass"`
	Price       int32     `json:"price"`
	OwnerData   OwnerData `json:"ownerData"`
	BookingID   string    `json:"bookingId"`
	FlightID    string    `json:"flightId"`
	UpdatedAt   string    `json:"updatedAt"`
}

type CancelTicketResponse struct {
	Message string        `json:"message"`
	Ticket  TicketDetails `json:"ticket"`
}

type GetTicketResponse struct {
	TicketID    string `json:"ticketId"`
	Status      string `json:"status"`
	SeatCode    string `json:"seatCode"`
	FlightClass string `json:"flightClass"`
	Price       int32  `json:"price"`
	OwnerData   struct {
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		Gender      string `json:"gender"`
		PhoneNumber string `json:"phoneNumber"`
	} `json:"ownerData"`
	BookingID string `json:"bookingId"`
	FlightID  string `json:"flightId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
