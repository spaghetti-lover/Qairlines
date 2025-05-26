package entities

import "time"

type TicketStatus string

const (
	TicketStatusBooked    TicketStatus = "booked"
	TicketStatusCancelled TicketStatus = "cancelled"
	TicketStatusCheckedIn TicketStatus = "checked_in"
	TicketStatusUsed      TicketStatus = "used"
)

type FlightClass string

const (
	FlightClassEconomy    FlightClass = "economy"
	FlightClassBusiness   FlightClass = "business"
	FlightClassFirstClass FlightClass = "first_class"
)

type Ticket struct {
	TicketID    string       `json:"ticket_id"`
	Status      TicketStatus `json:"status"`
	SeatCode    string       `json:"seat_code"`
	FlightClass FlightClass  `json:"flight_class"`
	Price       int32        `json:"price"`
	BookingID   int64        `json:"booking_id"`
	FlightID    int64        `json:"flight_id"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Owner       TicketOwner
	Seat        Seat
}
