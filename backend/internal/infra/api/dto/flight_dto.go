package dto

import (
	"time"

	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type CreateFlightRequest struct {
	FlightID         int64                 `json:"flightId"`
	FlightNumber     string                `json:"flightNumber"`
	AircraftType     string                `json:"aircraftType"`
	DepartureCity    string                `json:"departureCity"`
	ArrivalCity      string                `json:"arrivalCity"`
	DepartureAirport string                `json:"departureAirport"`
	ArrivalAirport   string                `json:"arrivalAirport"`
	DepartureTime    time.Time             `json:"departureTime"`
	ArrivalTime      time.Time             `json:"arrivalTime"`
	BasePrice        int32                 `json:"basePrice"`
	TotalSeatsRow    int32                 `json:"totalSeatsRow"`
	TotalSeatsColumn int32                 `json:"totalSeatsColumn"`
	Status           entities.FlightStatus `json:"status"`
}

type CreateFlightResponse struct {
	Message string `json:"message"`
	Flight  struct {
		FlightID         string `json:"flightId"`
		FlightNumber     string `json:"flightNumber"`
		AircraftType     string `json:"aircraftType"`
		DepartureCity    string `json:"departureCity"`
		ArrivalCity      string `json:"arrivalCity"`
		DepartureAirport string `json:"departureAirport"`
		ArrivalAirport   string `json:"arrivalAirport"`
		DepartureTime    string `json:"departureTime"`
		ArrivalTime      string `json:"arrivalTime"`
		BasePrice        int32  `json:"basePrice"`
		Status           string `json:"status"`
	} `json:"flight"`
}

type GetFlightResponse struct {
	FlightID      string `json:"flightId"`
	FlightNumber  string `json:"flightNumber"`
	AircraftType  string `json:"aircraftType"`
	DepartureCity string `json:"departureCity"`
	ArrivalCity   string `json:"arrivalCity"`
	DepartureTime struct {
		Seconds int64 `json:"seconds"`
	} `json:"departureTime"`
	ArrivalTime struct {
		Seconds int64 `json:"seconds"`
	} `json:"arrivalTime"`
	BasePrice int32                 `json:"basePrice"`
	Status    entities.FlightStatus `json:"status"`
}

type GetAllFlightsResponse struct {
	FlightID      string      `json:"flightId"`
	FlightNumber  string      `json:"flightNumber"`
	AircraftType  string      `json:"aircraftType"`
	DepartureCity string      `json:"departureCity"`
	ArrivalCity   string      `json:"arrivalCity"`
	DepartureTime TimeSeconds `json:"departureTime"`
	ArrivalTime   TimeSeconds `json:"arrivalTime"`
	BasePrice     int         `json:"basePrice"`
	Status        string      `json:"status"`
}

type FlightSearchResponse struct {
	FlightID         string `json:"flightId"`
	FlightNumber     string `json:"flightNumber"`
	Airline          string `json:"airline"`
	DepartureCity    string `json:"departureCity"`
	ArrivalCity      string `json:"arrivalCity"`
	DepartureTime    string `json:"departureTime"`
	ArrivalTime      string `json:"arrivalTime"`
	DepartureAirport string `json:"departureAirport"`
	ArrivalAirport   string `json:"arrivalAirport"`
	AircraftType     string `json:"aircraftType"`
	BasePrice        int    `json:"basePrice"`
}
