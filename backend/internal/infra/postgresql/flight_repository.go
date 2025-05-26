package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/spaghetti-lover/qairlines/db/sqlc"
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
