package entities

import "time"

type FlightStatus string

const (
	FlightOnTimeStatus   FlightStatus = "On Time"
	FlightDelayedStatus  FlightStatus = "Delayed"
	FlightCanceledStatus FlightStatus = "Cancelled"
	FlightBoardingStatus FlightStatus = "Boarding"
	FlightTakeoffStatus  FlightStatus = "Takeoff"
	FlightLandingStatus  FlightStatus = "Landing"
	FlightLandedStatus   FlightStatus = "Landed"
)

type Flight struct {
	FlightID         int64        `json:"flight_id"`
	FlightNumber     string       `json:"flight_number"`
	AircraftType     string       `json:"aircraft_type"`
	DepartureCity    string       `json:"departure_city"`
	ArrivalCity      string       `json:"arrival_city"`
	DepartureAirport string       `json:"departure_airport"`
	ArrivalAirport   string       `json:"arrival_airport"`
	DepartureTime    time.Time    `json:"departure_time"`
	ArrivalTime      time.Time    `json:"arrival_time"`
	BasePrice        int32        `json:"base_price"`
	TotalSeatsRow    int32        `json:"total_seats_row"`
	TotalSeatsColumn int32        `json:"total_seats_column"`
	Status           FlightStatus `json:"status"`
}

type CreateFlightParams struct {
	FlightID         string `json:"flight_id"`
	FlightNumber     string `json:"flight_number"`
	AircraftType     string `json:"aircraft_type"`
	DepartureCity    string `json:"departure_city"`
	ArrivalCity      string `json:"arrival_city"`
	DepartureAirport string `json:"departure_airport"`
	ArrivalAirport   string `json:"arrival_airport"`
	DepartureTime    string `json:"departure_time"`
	ArrivalTime      string `json:"arrival_time"`
	BasePrice        int32  `json:"base_price"`
	Status           string `json:"status"`
}
