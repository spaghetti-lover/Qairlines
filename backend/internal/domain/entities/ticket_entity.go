package entities

import "time"

type TicketStatus string

const (
	TicketBookedStatus    TicketStatus = "booked"
	TicketCancelledStatus TicketStatus = "cancelled"
	TicketUsedStatus      TicketStatus = "used"
)

type Ticket struct {
	TicketID    string       `json:"ticket_id"`
	Status      TicketStatus `json:"status"`
	FlightClass FlightStatus `json:"flight_class"`
	Price       int32        `json:"price"`
	BookingID   string       `json:"booking_id"`
	FlightID    string       `json:"flight_id"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Owner       TicketOwner
	Seat        Seat
}
