package entities

import (
	"time"
)

type TicketStatus string

const (
	TicketStatusActive    TicketStatus = "Booked"
	TicketStatusCancelled TicketStatus = "Cancelled"
)

type FlightClass string

const (
	FlightClassEconomy    FlightClass = "economy"
	FlightClassBusiness   FlightClass = "business"
	FlightClassFirstClass FlightClass = "firstClass"
)

type Ticket struct {
	TicketID    int64        `json:"ticket_id"`
	SeatID      int64        `json:"seat_id"`
	FlightClass FlightClass  `json:"flight_class"`
	Price       int32        `json:"price"`
	Status      TicketStatus `json:"status"`
	BookingID   int64        `json:"booking_id"`
	FlightID    int64        `json:"flight_id"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Seat        Seat         `json:"seat"`
	Owner       TicketOwner  `json:"owner"`
}
