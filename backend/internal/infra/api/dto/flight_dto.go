package dto

import (
	"time"

	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type CreateFlightRequest struct {
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
type GetFlightsWithTicketsResponse struct {
	Flights []FlightWithTickets `json:"flights"`
}

type FlightWithTickets struct {
	FlightID      string           `json:"flight_id"`
	FlightNumber  string           `json:"flight_number"`
	AircraftType  string           `json:"aircraft_type"`
	DepartureCity string           `json:"departure_city"`
	ArrivalCity   string           `json:"arrival_city"`
	DepartureTime TimeSeconds      `json:"departure_time"`
	ArrivalTime   TimeSeconds      `json:"arrival_time"`
	BasePrice     int              `json:"base_price"`
	Status        string           `json:"status"`
	TicketList    []TicketResponse `json:"ticket_list"`
}

type TicketResponse struct {
	TicketID    int64  `json:"ticket_id"`
	SeatCode    string `json:"seat_code"`
	Price       int32  `json:"price"`
	FlightClass string `json:"flight_class"`
	Status      string `json:"status"`
}

type ListFlightsParams struct {
	Limit int `json:"limit" binding:"required,min=1,max=100" default:"10"`
	Page  int `json:"page" binding:"required,min=1" default:"1"`
}
