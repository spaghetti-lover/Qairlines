package entities

type BookingStatus string

const (
	BookingConfirmedStatus BookingStatus = "confirmed"
	BookingCancelledStatus BookingStatus = "cancelled"
	BookingPendingStatus   BookingStatus = "pending"
)

type TripType string

const (
	OneWayTrip TripType = "one_way"
	RoundTrip  TripType = "round_trip"
)

type Booking struct {
	BookingID         string        `json:"booking_id"`
	UserEmail         string        `json:"user_email"`
	TripType          TripType      `json:"trip_type"`
	DepartureFlightID string        `json:"departure_flight_id"`
	ReturnFlightID    string        `json:"return_flight_id"`
	Status            BookingStatus `json:"status"`
	CreatedAt         string        `json:"created_at"`
	UpdatedAt         string        `json:"updated_at"`
}
type CreateBookingParams struct {
	BookingID         string        `json:"booking_id"`
	UserEmail         string        `json:"user_email"`
	TripType          TripType      `json:"trip_type"`
	DepartureFlightID string        `json:"departure_flight_id"`
	ReturnFlightID    string        `json:"return_flight_id"`
	Status            BookingStatus `json:"status"`
}
