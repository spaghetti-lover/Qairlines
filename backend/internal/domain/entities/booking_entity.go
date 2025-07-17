package entities

import "time"

type BookingStatus string

const (
	BookingConfirmedStatus BookingStatus = "confirmed"
	BookingCancelledStatus BookingStatus = "cancelled"
	BookingPendingStatus   BookingStatus = "pending"
)

type TripType string

const (
	OneWayTrip TripType = "oneWay"
	RoundTrip  TripType = "roundTrip"
)

const (
	BookingStatusConfirmed BookingStatus = "confirmed"
	BookingStatusCancelled BookingStatus = "cancelled"
	BookingStatusPending   BookingStatus = "pending"
)

type Booking struct {
	BookingID         int64         `json:"booking_id"`
	UserEmail         string        `json:"user_email"`
	TripType          TripType      `json:"trip_type"`
	DepartureFlightID int64         `json:"departure_flight_id"`
	ReturnFlightID    *int64        `json:"return_flight_id,omitempty"`
	Status            BookingStatus `json:"status"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
}

type CreateBookingParams struct {
	Email                   string   `json:"email"`
	DepartureCity           string   `json:"departureCity"`
	ArrivalCity             string   `json:"arrivalCity"`
	DepartureFlightID       string   `json:"departureFlightId"`
	ReturnFlightID          string   `json:"returnFlightId"`
	TripType                TripType `json:"tripType"`
	DepartureTicketDataList []Ticket `json:"departureTicketDataList"`
	ReturnTicketDataList    []Ticket `json:"returnTicketDataList"`
	AfterCreate             func(booking Booking) error
}
