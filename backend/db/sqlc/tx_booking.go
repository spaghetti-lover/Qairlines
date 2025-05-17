package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/spaghetti-lover/qairlines/pkg/utils"
)

// BookingTxParams contains the input parameters of the booking transaction
type BookingTxParams struct {
	BookerEmail      string          `json:"booker_email"`
	NumberOfAdults   int64           `json:"number_of_adults"`
	NumberOfChildren int64           `json:"number_of_children"`
	FlightClass      FlightClassType `json:"flight_class"`
	Cancelled        pgtype.Bool     `json:"cancelled"`
	FlightID         int64           `json:"flight_id"`
}

// BookingTxResult is the result of the booking transaction
type BookingTxResult struct {
	Booking          Booking         `json:"booking"`
	BookingID        string          `json:"booking_id"`
	BookerEmail      string          `json:"booker_email"`
	NumberOfAdults   int64           `json:"number_of_adults"`
	NumberOfChildren int64           `json:"number_of_children"`
	FlightClass      FlightClassType `json:"flight_class"`
	Cancelled        pgtype.Bool     `json:"cancelled"`
	FlightID         int64           `json:"flight_id"`
}

type PassengerParams struct {
	BookingID      int64       `json:"booking_id"`
	CitizenID      string      `json:"citizen_id"`
	PassportNumber pgtype.Text `json:"passport_number"`
	Gender         GenderEnum  `json:"gender"`
	PhoneNumber    string      `json:"phone_number"`
	FirstName      string      `json:"first_name"`
	LastName       string      `json:"last_name"`
	Nationality    string      `json:"nationality"`
	DateOfBirth    pgtype.Date `json:"date_of_birth"`
	SeatRow        int32       `json:"seat_row"`
	SeatCol        string      `json:"seat_col"`
}

type PaymentParams struct {
	Amount        pgtype.Numeric `json:"amount"`
	Currency      pgtype.Text    `json:"currency"`
	PaymentMethod pgtype.Text    `json:"payment_method"`
	Status        pgtype.Text    `json:"status"`
	BookingID     int64          `json:"booking_id"`
}

// BookingTx performs a booking transfer from user
func (store *SQLStore) BookingTx(ctx context.Context, arg BookingTxParams, passengers []PassengerParams, payment PaymentParams) (BookingTxResult, error) {
	var result BookingTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		// Create booking query
		result.Booking, err = q.CreateBooking(ctx, CreateBookingParams(arg))
		if err != nil {
			return err
		}

		// Encode booking_id
		encodedID, err := utils.EncodeBookingID(result.Booking.BookingID)
		if err != nil {
			return err
		}

		// Ensure no overbooking
		flightSeat, err := q.GetFlightSeat(ctx, GetFlightSeatParams{
			FlightID:    arg.FlightID,
			FlightClass: arg.FlightClass,
		})

		if err != nil {
			return fmt.Errorf("can not find flight_seat info: %w", err)
		}
		maxSeats := flightSeat.MaxRowSeat * flightSeat.MaxColSeat
		occupied, err := q.CountOccupiedSeats(ctx, CountOccupiedSeatsParams{
			FlightID:    arg.FlightID,
			FlightClass: arg.FlightClass,
		})
		if err != nil {
			return fmt.Errorf("can not count occupied seats: %w", err)
		}
		if int(occupied)+len(passengers) > int(maxSeats) {
			return fmt.Errorf("exceed seat limit for class %s: occupied %d, add %d, sum %d > %d",
				arg.FlightClass, occupied, len(passengers), int(occupied)+len(passengers), maxSeats)
		}

		// Check valid seat and not occupied
		for _, p := range passengers {
			if err := validateSeat(p, flightSeat); err != nil {
				return err
			}
			if err := checkSeatAvailability(ctx, q, arg, p); err != nil {
				return err
			}
		}

		// Update passenger info
		if err := createPassengers(ctx, q, passengers, result.Booking.BookingID); err != nil {
			return err
		}
		// TODO: Add payment info

		// Gán ghế ngồi cho hành khách.
		result = BookingTxResult{
			Booking:          result.Booking,
			BookingID:        encodedID,
			BookerEmail:      result.Booking.BookerEmail,
			NumberOfAdults:   result.Booking.NumberOfAdults,
			NumberOfChildren: result.Booking.NumberOfChildren,
			FlightClass:      result.Booking.FlightClass,
			Cancelled:        result.Booking.Cancelled,
			FlightID:         result.Booking.FlightID,
		}
		return err
	})
	return result, err
}

func validateSeat(p PassengerParams, seat FlightSeat) error {
	if len(p.SeatCol) != 1 || rune(p.SeatCol[0]) < 'A' || rune(p.SeatCol[0]) >= rune('A'+seat.MaxColSeat) {
		return fmt.Errorf("invalid seat column: %s with maxcol is: %v, maxrow is: %v", p.SeatCol, seat.MaxColSeat, seat.MaxRowSeat)
	}
	if p.SeatRow <= 0 || p.SeatRow > int32(seat.MaxRowSeat) {
		return fmt.Errorf("invalid seat row: %d", p.SeatRow)
	}
	return nil
}

func checkSeatAvailability(ctx context.Context, q *Queries, arg BookingTxParams, p PassengerParams) error {
	isOccupied, err := q.CheckSeatOccupied(ctx, CheckSeatOccupiedParams{
		FlightID:    arg.FlightID,
		FlightClass: arg.FlightClass,
		SeatRow:     p.SeatRow,
		SeatCol:     p.SeatCol,
	})
	if err != nil {
		return fmt.Errorf("error checking seat: %w", err)
	}
	if isOccupied {
		return fmt.Errorf("seat %d%s is already booked", p.SeatRow, p.SeatCol)
	}
	return nil
}

func createPassengers(ctx context.Context, q *Queries, passengers []PassengerParams, bookingID int64) error {
	for _, p := range passengers {
		p.BookingID = bookingID
		_, err := q.CreatePassenger(ctx, CreatePassengerParams(p))
		if err != nil {
			return err
		}
	}
	return nil
}
