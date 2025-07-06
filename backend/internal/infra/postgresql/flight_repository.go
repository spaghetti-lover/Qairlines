package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	db "github.com/spaghetti-lover/qairlines/db/sqlc"
	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type FlightRepositoryPostgres struct {
	store db.Store
}

func NewFlightRepositoryPostgres(store *db.Store) *FlightRepositoryPostgres {
	return &FlightRepositoryPostgres{store: *store}
}

func (r *FlightRepositoryPostgres) CreateFlight(ctx context.Context, flight entities.Flight) (entities.Flight, error) {
	dbFlight, err := r.store.CreateFlight(ctx, db.CreateFlightParams{
		FlightNumber:     flight.FlightNumber,
		AircraftType:     &flight.AircraftType,
		DepartureCity:    &flight.DepartureCity,
		ArrivalCity:      &flight.ArrivalCity,
		DepartureAirport: &flight.DepartureAirport,
		ArrivalAirport:   &flight.ArrivalAirport,
		DepartureTime:    flight.DepartureTime,
		ArrivalTime:      flight.ArrivalTime,
		BasePrice:        flight.BasePrice,
		Status:           db.FlightStatus(flight.Status),
	})
	if err != nil {
		return entities.Flight{}, err
	}

	return entities.Flight{
		FlightNumber:     dbFlight.FlightNumber,
		AircraftType:     *dbFlight.AircraftType,
		DepartureCity:    *dbFlight.DepartureCity,
		ArrivalCity:      *dbFlight.ArrivalCity,
		DepartureAirport: *dbFlight.DepartureAirport,
		ArrivalAirport:   *dbFlight.ArrivalAirport,
		DepartureTime:    dbFlight.DepartureTime,
		ArrivalTime:      dbFlight.ArrivalTime,
		BasePrice:        dbFlight.BasePrice,
		Status:           entities.FlightStatus(dbFlight.Status),
	}, nil
}

func (r *FlightRepositoryPostgres) GetFlightByID(ctx context.Context, flightID int64) (*entities.Flight, error) {
	dbFlight, err := r.store.GetFlight(ctx, flightID)
	if err != nil {
		return nil, err
	}
	return &entities.Flight{
		FlightID:         dbFlight.FlightID,
		FlightNumber:     dbFlight.FlightNumber,
		AircraftType:     *dbFlight.AircraftType,
		DepartureCity:    *dbFlight.DepartureCity,
		ArrivalCity:      *dbFlight.ArrivalCity,
		DepartureAirport: *dbFlight.DepartureAirport,
		ArrivalAirport:   *dbFlight.ArrivalAirport,
		DepartureTime:    dbFlight.DepartureTime,
		ArrivalTime:      dbFlight.ArrivalTime,
		BasePrice:        dbFlight.BasePrice,
		Status:           entities.FlightStatus(dbFlight.Status),
	}, nil
}

func (r *FlightRepositoryPostgres) UpdateFlightTimes(ctx context.Context, flightID int64, departureTime, arrivalTime time.Time) (*entities.Flight, error) {
	row, err := r.store.UpdateFlightTimes(ctx, db.UpdateFlightTimesParams{
		FlightID:      flightID,
		DepartureTime: departureTime,
		ArrivalTime:   arrivalTime,
	})
	if err == sql.ErrNoRows {
		return nil, adapters.ErrFlightNotFound
	}
	if err != nil {
		return nil, err
	}

	return &entities.Flight{
		FlightID:      row.FlightID,
		DepartureTime: row.DepartureTime,
		ArrivalTime:   row.ArrivalTime,
	}, nil
}

func (r *FlightRepositoryPostgres) GetAllFlights(ctx context.Context) ([]entities.Flight, error) {
	rows, err := r.store.GetAllFlights(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all flights: %w", err)
	}

	var flights []entities.Flight
	for _, row := range rows {
		flights = append(flights, entities.Flight{
			FlightID:      row.FlightID,
			FlightNumber:  row.FlightNumber,
			AircraftType:  *row.AircraftType,
			DepartureCity: *row.DepartureCity,
			ArrivalCity:   *row.ArrivalCity,
			DepartureTime: row.DepartureTime,
			ArrivalTime:   row.ArrivalTime,
			BasePrice:     row.BasePrice,
			Status:        entities.FlightStatus(row.Status),
		})
	}

	return flights, nil
}

func (r *FlightRepositoryPostgres) DeleteFlightByID(ctx context.Context, flightID int64) error {
	flighID, err := r.store.DeleteFlight(ctx, flightID)
	if flighID == 0 {
		return adapters.ErrFlightNotFound
	}
	if err != nil {
		return fmt.Errorf("failed to delete flight: %w", err)
	}

	return nil
}

func (r *FlightRepositoryPostgres) SearchFlights(ctx context.Context, departureCity, arrivalCity string, flightDate time.Time) ([]entities.Flight, error) {
	rows, err := r.store.SearchFlights(ctx, db.SearchFlightsParams{
		DepartureCity: &departureCity,
		ArrivalCity:   &arrivalCity,
		DepartureTime: flightDate,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to search flights: %w", err)
	}

	if len(rows) == 0 {
		return nil, adapters.ErrNoFlightsFound
	}

	var flights []entities.Flight
	for _, row := range rows {
		flights = append(flights, entities.Flight{
			FlightID:         row.FlightID,
			FlightNumber:     row.FlightNumber,
			Airline:          *row.Airline,
			DepartureCity:    *row.DepartureCity,
			ArrivalCity:      *row.ArrivalCity,
			DepartureTime:    row.DepartureTime,
			ArrivalTime:      row.ArrivalTime,
			DepartureAirport: *row.DepartureAirport,
			ArrivalAirport:   *row.ArrivalAirport,
			AircraftType:     *row.AircraftType,
			BasePrice:        row.BasePrice,
		})
	}

	return flights, nil
}

func (r *FlightRepositoryPostgres) GetSuggestedFlights(ctx context.Context) ([]entities.Flight, error) {
	rows, err := r.store.GetSuggestedFlights(ctx)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, adapters.ErrNoSuggestedFlights
	}

	var flights []entities.Flight
	for _, row := range rows {
		flights = append(flights, entities.Flight{
			FlightID:         row.FlightID,
			FlightNumber:     row.FlightNumber,
			Airline:          *row.Airline,
			DepartureCity:    *row.DepartureCity,
			ArrivalCity:      *row.ArrivalCity,
			DepartureTime:    row.DepartureTime,
			ArrivalTime:      row.ArrivalTime,
			DepartureAirport: *row.DepartureAirport,
			ArrivalAirport:   *row.ArrivalAirport,
			AircraftType:     *row.AircraftType,
			BasePrice:        row.BasePrice,
		})
	}

	return flights, nil
}
