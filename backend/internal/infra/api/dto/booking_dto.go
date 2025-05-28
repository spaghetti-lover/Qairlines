package dto

import (
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type CreateBookingRequest struct {
	DepartureCity           string              `json:"departureCity"`
	ArrivalCity             string              `json:"arrivalCity"`
	DepartureFlightID       string              `json:"departureFlightId"`
	ReturnFlightID          string              `json:"returnFlightId"`
	TripType                string              `json:"tripType"`
	DepartureTicketDataList []TicketDataRequest `json:"departureTicketDataList"`
	ReturnTicketDataList    []TicketDataRequest `json:"returnTicketDataList"`
}

type TicketDataRequest struct {
	Price       int32     `json:"price"`
	FlightClass string    `json:"flightClass"`
	OwnerData   OwnerData `json:"ownerData"`
}

type OwnerData struct {
	IdentityCardNumber string              `json:"identityCardNumber"`
	FirstName          string              `json:"firstName"`
	LastName           string              `json:"lastName"`
	PhoneNumber        string              `json:"phoneNumber"`
	DateOfBirth        string              `json:"dateOfBirth"`
	Gender             entities.GenderType `json:"gender"`
	Address            string              `json:"address"`
}

type CreateBookingResponse struct {
	BookingID         string               `json:"bookingId"`
	DepartureFlightID string               `json:"departureFlightId"`
	ReturnFlightID    string               `json:"returnFlightId"`
	TripType          string               `json:"tripType"`
	DepartureTickets  []TicketDataResponse `json:"departureTickets"`
	ReturnTickets     []TicketDataResponse `json:"returnTickets"`
}

type TicketDataResponse struct {
	TicketID    string    `json:"ticketId"`
	SeatID      *string   `json:"seatId,omitempty"`
	Price       int32     `json:"price"`
	FlightClass string    `json:"flightClass"`
	OwnerData   OwnerData `json:"ownerData"`
}
