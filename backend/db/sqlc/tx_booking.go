package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type BookingTxParams struct {
	BookingID         int64
	UserEmail         string
	TripType          string
	DepartureFlightID int64
	ReturnFlightID    int64
	Status            string
	CreatedAt         pgtype.Timestamptz
	UpdatedAt         pgtype.Timestamptz
}

type BookingTxResult struct {
	BookingID         int64
	UserEmail         string
	TripType          string
	DepartureFlightID int64
	ReturnFlightID    int64
	Status            string
	CreatedAt         pgtype.Timestamptz
	UpdatedAt         pgtype.Timestamptz
}

// BookingTx performs a booking transfer from user
func (store *SQLStore) BookingTx(ctx context.Context, arg BookingTxParams) (BookingTxResult, error) {
	var result BookingTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		// Check user_email exists
		_, err := store.GetUserByEmail(ctx, arg.UserEmail)
		if err != nil {
			return fmt.Errorf("user with email %s does not exist: %w", arg.UserEmail, err)
		}
		// Check departure flight id, return flight id exists
		_, err = store.GetFlight(ctx, arg.DepartureFlightID)
		if err != nil {
			return fmt.Errorf("departure flight with id %d does not exist: %w", arg.DepartureFlightID, err)
		}
		_, err = store.GetFlight(ctx, arg.ReturnFlightID)
		if err != nil && arg.TripType == "round_trip" {
			return fmt.Errorf("return flight with id %d does not exist: %w", arg.ReturnFlightID, err)
		}
		// Check trip type is valid. If trip type is round trip, return flight id must be not null
		if arg.TripType != "one_way" && arg.TripType != "round_trip" {
			return fmt.Errorf("invalid trip type: %s", arg.TripType)
		}
		if arg.TripType == "round_trip" && arg.ReturnFlightID <= 0 {
			return fmt.Errorf("return flight id must be not null for round trip")
		}
		// Check flight status is valid
		flightStatus, err := store.GetFlightsByStatus(ctx, arg.DepartureFlightID)
		if err != nil {
			return fmt.Errorf("failed to get flight status for departure flight %d: %w", arg.DepartureFlightID, err)
		}
		if flightStatus != "On Time" && flightStatus != "Boarding" {
			return fmt.Errorf("departure flight with id %d is not valid for booking: %s", arg.DepartureFlightID, flightStatus)
		}

		// Check enough seats available for the trip type
		availableDepartureSeats, err := store.CountOccupiedSeats(ctx, pgtype.Int8{Int64: arg.DepartureFlightID, Valid: true})
		if err != nil {
			return fmt.Errorf("failed to count occupied seats for departure flight %d: %w", arg.DepartureFlightID, err)
		}
		availableReturnSeats := int64(0)
		if arg.TripType == "round_trip" {
			availableReturnSeats, err = store.CountOccupiedSeats(ctx, pgtype.Int8{Int64: arg.ReturnFlightID, Valid: true})
			if err != nil {
				return fmt.Errorf("failed to count occupied seats for return flight %d: %w", arg.ReturnFlightID, err)
			}
		}
		if availableDepartureSeats <= 0 {
			return fmt.Errorf("no available seats for departure flight %d", arg.DepartureFlightID)
		}
		if arg.TripType == "round_trip" && availableReturnSeats <= 0 {
			return fmt.Errorf("no available seats for return flight %d", arg.ReturnFlightID)
		}

		// Create booking
		bookingParams := CreateBookingParams{
			BookingID:         arg.BookingID,
			UserEmail:         pgtype.Text{String: arg.UserEmail, Valid: true},
			TripType:          TripType(arg.TripType),
			DepartureFlightID: arg.DepartureFlightID,
			ReturnFlightID:    pgtype.Int8{Int64: arg.ReturnFlightID, Valid: arg.TripType == "round_trip"},
			Status:            BookingStatus("pending"),
		}
		booking, err := store.CreateBooking(ctx, bookingParams)
		if err != nil {
			return fmt.Errorf("failed to create booking: %w", err)
		}
		// Create Ticket
		ticketParams := CreateTicketParams{
			BookingID:   pgtype.Int8{Int64: booking.BookingID, Valid: true},
			FlightID:    pgtype.Int8{Int64: arg.DepartureFlightID, Valid: true},
			FlightClass: FlightClass("Economy"), // This should be determined based on the booking or flight
			Price:       100000,                 // This should be determined based on the flight and class
			Status:      TicketStatus("pending"),
		}
		_, err = store.CreateTicket(ctx, ticketParams)
		if err != nil {
			return fmt.Errorf("failed to create ticket: %w", err)
		}
		// Result
		result = BookingTxResult{
			BookingID:         booking.BookingID,
			UserEmail:         booking.UserEmail.String,
			TripType:          string(booking.TripType),
			DepartureFlightID: booking.DepartureFlightID,
			ReturnFlightID:    booking.ReturnFlightID.Int64,
			Status:            string(booking.Status),
			CreatedAt:         pgtype.Timestamptz{Time: booking.CreatedAt, Valid: true},
			UpdatedAt:         pgtype.Timestamptz{Time: booking.UpdatedAt, Valid: true},
		}
		return err
	})
	return result, err
}
