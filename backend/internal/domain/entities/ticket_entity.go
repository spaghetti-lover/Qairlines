package entities

type TicketStatus string

const (
	TicketBookedStatus    TicketStatus = "booked"
	TicketCancelledStatus TicketStatus = "cancelled"
	TicketUsedStatus      TicketStatus = "used"
)

type Ticket struct {
	TicketID    string `json:"ticket_id"`
	SeatCode    string `json:"seat_code"`
	FlightClass string `json:"flight_class"`
	Price       int32  `json:"price"`
	Status      string `json:"status"`
	BookingID   string `json:"booking_id"`
	FlightID    string `json:"flight_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
