package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
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
		FlightID:         flight.FlightID,
		FlightNumber:     flight.FlightNumber,
		AircraftType:     pgtype.Text{String: flight.AircraftType, Valid: true},
		DepartureCity:    pgtype.Text{String: flight.DepartureCity, Valid: true},
		ArrivalCity:      pgtype.Text{String: flight.ArrivalCity, Valid: true},
		DepartureAirport: pgtype.Text{String: flight.DepartureAirport, Valid: true},
		ArrivalAirport:   pgtype.Text{String: flight.ArrivalAirport, Valid: true},
		DepartureTime:    flight.DepartureTime,
		ArrivalTime:      flight.ArrivalTime,
		BasePrice:        flight.BasePrice,
		Status:           db.FlightStatus(flight.Status),
	})
	if err != nil {
		return entities.Flight{}, err
	}

	return entities.Flight{
		FlightID:         dbFlight.FlightID,
		FlightNumber:     dbFlight.FlightNumber,
		AircraftType:     dbFlight.AircraftType.String,
		DepartureCity:    dbFlight.DepartureCity.String,
		ArrivalCity:      dbFlight.ArrivalCity.String,
		DepartureAirport: dbFlight.DepartureAirport.String,
		ArrivalAirport:   dbFlight.ArrivalAirport.String,
		DepartureTime:    dbFlight.DepartureTime,
		ArrivalTime:      dbFlight.ArrivalTime,
		BasePrice:        dbFlight.BasePrice,
		Status:           entities.FlightStatus(dbFlight.Status),
	}, nil
}

func (r *FlightRepositoryPostgres) GetFlightByID(ctx context.Context, flightID int64) (*entities.Flight, error) {
	dbFlight, err := r.store.GetFlight(ctx, flightID)
	if err != nil {
		return nil, adapters.ErrFlightNotFound
	}
	return &entities.Flight{
		FlightID:         dbFlight.FlightID,
		FlightNumber:     dbFlight.FlightNumber,
		AircraftType:     dbFlight.AircraftType.String,
		DepartureCity:    dbFlight.DepartureCity.String,
		ArrivalCity:      dbFlight.ArrivalCity.String,
		DepartureAirport: dbFlight.DepartureAirport.String,
		ArrivalAirport:   dbFlight.ArrivalAirport.String,
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
			AircraftType:  row.AircraftType.String,
			DepartureCity: row.DepartureCity.String,
			ArrivalCity:   row.ArrivalCity.String,
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
		DepartureCity: pgtype.Text{String: departureCity, Valid: true},
		ArrivalCity:   pgtype.Text{String: arrivalCity, Valid: true},
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
			Airline:          row.Airline.String,
			DepartureCity:    row.DepartureCity.String,
			ArrivalCity:      row.ArrivalCity.String,
			DepartureTime:    row.DepartureTime,
			ArrivalTime:      row.ArrivalTime,
			DepartureAirport: row.DepartureAirport.String,
			ArrivalAirport:   row.ArrivalAirport.String,
			AircraftType:     row.AircraftType.String,
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
			Airline:          row.Airline.String,
			DepartureCity:    row.DepartureCity.String,
			ArrivalCity:      row.ArrivalCity.String,
			DepartureTime:    row.DepartureTime,
			ArrivalTime:      row.ArrivalTime,
			DepartureAirport: row.DepartureAirport.String,
			ArrivalAirport:   row.ArrivalAirport.String,
			AircraftType:     row.AircraftType.String,
			BasePrice:        row.BasePrice,
		})
	}

	return flights, nil
}
