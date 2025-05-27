package mappers

import (
	"strconv"

	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
)

func CreateFlightRequestToEntity(req dto.CreateFlightRequest) entities.Flight {
	return entities.Flight{
		FlightID:         req.FlightID,
		FlightNumber:     req.FlightNumber,
		AircraftType:     req.AircraftType,
		DepartureCity:    req.DepartureCity,
		ArrivalCity:      req.ArrivalCity,
		DepartureAirport: req.DepartureAirport,
		ArrivalAirport:   req.ArrivalAirport,
		DepartureTime:    req.DepartureTime,
		ArrivalTime:      req.ArrivalTime,
		BasePrice:        req.BasePrice,
		Status:           req.Status,
	}
}

func CreateFlightEntityToResponse(flight entities.Flight) dto.CreateFlightResponse {
	return dto.CreateFlightResponse{
		Message: "Flight created successfully.",
		Flight: struct {
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
		}{
			FlightID:         strconv.FormatInt(flight.FlightID, 10),
			FlightNumber:     flight.FlightNumber,
			AircraftType:     flight.AircraftType,
			DepartureCity:    flight.DepartureCity,
			ArrivalCity:      flight.ArrivalCity,
			DepartureAirport: flight.DepartureAirport,
			ArrivalAirport:   flight.ArrivalAirport,
			DepartureTime:    flight.DepartureTime.String(),
			ArrivalTime:      flight.ArrivalTime.String(),
			BasePrice:        flight.BasePrice,
			Status:           string(flight.Status),
		},
	}
}

func MapFlightToGetFlightResponse(flight *entities.Flight) *dto.GetFlightResponse {
	return &dto.GetFlightResponse{
		FlightID:      strconv.FormatInt(flight.FlightID, 10),
		FlightNumber:  flight.FlightNumber,
		AircraftType:  flight.AircraftType,
		DepartureCity: flight.DepartureCity,
		ArrivalCity:   flight.ArrivalCity,
		DepartureTime: struct {
			Seconds int64 `json:"seconds"`
		}{Seconds: flight.DepartureTime.Unix()},
		ArrivalTime: struct {
			Seconds int64 `json:"seconds"`
		}{Seconds: flight.ArrivalTime.Unix()},
		BasePrice: flight.BasePrice,
		Status:    entities.FlightStatus(flight.Status),
	}
}

func ToUpdateFlightTimesResponse(flight *entities.Flight) *dto.UpdateFlightTimesResponse {
	return &dto.UpdateFlightTimesResponse{
		FlightID: flight.FlightID,
		DepartureTime: dto.TimeSeconds{
			Seconds: flight.DepartureTime.Unix(),
		},
		ArrivalTime: dto.TimeSeconds{
			Seconds: flight.ArrivalTime.Unix(),
		},
	}
}

func ToFlightResponses(flights []entities.Flight) []dto.GetAllFlightsResponse {
	var flightResponses []dto.GetAllFlightsResponse
	for _, flight := range flights {
		flightResponses = append(flightResponses, dto.GetAllFlightsResponse{
			FlightID:      strconv.FormatInt(flight.FlightID, 10),
			FlightNumber:  flight.FlightNumber,
			AircraftType:  flight.AircraftType,
			DepartureCity: flight.DepartureCity,
			ArrivalCity:   flight.ArrivalCity,
			DepartureTime: dto.TimeSeconds{Seconds: flight.DepartureTime.Unix()},
			ArrivalTime:   dto.TimeSeconds{Seconds: flight.ArrivalTime.Unix()},
			BasePrice:     int(flight.BasePrice),
			Status:        string(flight.Status),
		})
	}
	return flightResponses
}

func ToFlightSearchResponses(flights []entities.Flight) []dto.FlightSearchResponse {
	var responses []dto.FlightSearchResponse

	for _, flight := range flights {
		responses = append(responses, dto.FlightSearchResponse{
			FlightID:         strconv.FormatInt(flight.FlightID, 10),
			FlightNumber:     flight.FlightNumber,
			Airline:          flight.Airline,
			DepartureCity:    flight.DepartureCity,
			ArrivalCity:      flight.ArrivalCity,
			DepartureTime:    flight.DepartureTime.Format("2006-01-02T15:04:05Z"),
			ArrivalTime:      flight.ArrivalTime.Format("2006-01-02T15:04:05Z"),
			DepartureAirport: flight.DepartureAirport,
			ArrivalAirport:   flight.ArrivalAirport,
			AircraftType:     flight.AircraftType,
			BasePrice:        int(flight.BasePrice),
		})
	}

	return responses
}
