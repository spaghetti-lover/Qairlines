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
