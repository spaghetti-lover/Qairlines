package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type CreateBookingTxParams struct {
	UserEmail           string
	DepartureCity       string
	ArrivalCity         string
	DepartureFlightID   int64
	ReturnFlightID      int64
	TripType            string
	DepartureTicketData []TicketData
	ReturnTicketData    []TicketData
}

type TicketData struct {
	Price       int64
	FlightClass string
	OwnerData   OwnerData
}

type OwnerData struct {
	IdentityCardNumber string
	FirstName          string
	LastName           string
	PhoneNumber        string
	DateOfBirth        string
	Gender             string
	Address            string
}

type CreateBookingTxResult struct {
	Booking          entities.Booking
	DepartureTickets []entities.Ticket
	ReturnTickets    []entities.Ticket
}

func (store *SQLStore) CreateBookingTx(ctx context.Context, arg CreateBookingTxParams) (CreateBookingTxResult, error) {
	var result CreateBookingTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		// Tạo booking
		var returnFlightIDPtr *int64
		if arg.TripType == "roundTrip" && arg.ReturnFlightID != 0 {
			returnFlightIDPtr = &arg.ReturnFlightID
		} else {
			returnFlightIDPtr = nil
		}

		booking, err := q.CreateBooking(ctx, CreateBookingParams{
			UserEmail:         pgtype.Text{String: arg.UserEmail, Valid: true},
			TripType:          TripType(arg.TripType),
			DepartureFlightID: pgtype.Int8{Int64: arg.DepartureFlightID, Valid: true},
			ReturnFlightID:    pgtype.Int8{Int64: *returnFlightIDPtr, Valid: returnFlightIDPtr != nil},
			Status:            BookingStatus(entities.BookingStatusPending),
		})
		if err != nil {
			return fmt.Errorf("failed to create booking: %w", err)
		}
		result.Booking = entities.Booking{
			BookingID:         booking.BookingID,
			UserEmail:         booking.UserEmail.String,
			TripType:          entities.TripType(booking.TripType),
			DepartureFlightID: *&booking.DepartureFlightID.Int64,
			ReturnFlightID:    &booking.ReturnFlightID.Int64,
			Status:            entities.BookingStatus(booking.Status),
			CreatedAt:         booking.CreatedAt,
			UpdatedAt:         booking.UpdatedAt,
		}

		// Tạo vé cho chuyến bay đi
		for _, ticket := range arg.DepartureTicketData {
			createdTicket, err := createTicketForBooking(ctx, q, booking.BookingID, arg.DepartureFlightID, ticket)
			if err != nil {
				return err
			}
			result.DepartureTickets = append(result.DepartureTickets, createdTicket)
		}

		// Tạo vé cho chuyến bay về (nếu có)
		if arg.TripType == "roundTrip" && arg.ReturnFlightID != 0 {
			for _, ticket := range arg.ReturnTicketData {
				createdTicket, err := createTicketForBooking(ctx, q, booking.BookingID, arg.ReturnFlightID, ticket)
				if err != nil {
					return err
				}
				result.ReturnTickets = append(result.ReturnTickets, createdTicket)
			}
		}

		return nil
	})

	return result, err
}

func createTicketForBooking(ctx context.Context, q *Queries, bookingID int64, flightID int64, ticket TicketData) (entities.Ticket, error) {
	createdSeat, err := q.CreateSeat(ctx, CreateSeatParams{
		IsAvailable: true,
		Class:       FlightClass(ticket.FlightClass),
		FlightID:    pgtype.Int8{Int64: flightID, Valid: true},
	})

	if err != nil {
		return entities.Ticket{}, fmt.Errorf("failed to create seat: %w", err)
	}

	createdTicket, err := q.CreateTicket(ctx, CreateTicketParams{
		SeatID:      createdSeat.SeatID,
		FlightClass: FlightClass(ticket.FlightClass),
		Price:       int32(ticket.Price),
		Status:      TicketStatusActive,
		BookingID:   pgtype.Int8{Int64: bookingID, Valid: true},
		FlightID:    flightID,
	})

	if err != nil {
		return entities.Ticket{}, fmt.Errorf("failed to create ticket: %w", err)
	}

	_, err = q.CreateTicketOwnerSnapshot(ctx, CreateTicketOwnerSnapshotParams{
		TicketID:       createdTicket.TicketID,
		FirstName:      pgtype.Text{String: ticket.OwnerData.FirstName, Valid: true},
		LastName:       pgtype.Text{String: ticket.OwnerData.LastName, Valid: true},
		PhoneNumber:    pgtype.Text{String: ticket.OwnerData.PhoneNumber, Valid: true},
		Gender:         GenderType(ticket.OwnerData.Gender),
		DateOfBirth:    parseDate(ticket.OwnerData.DateOfBirth),
		PassportNumber: pgtype.Text{String: ticket.OwnerData.IdentityCardNumber, Valid: true},
		Address:        pgtype.Text{String: ticket.OwnerData.Address, Valid: true},
	})
	if err != nil {
		return entities.Ticket{}, fmt.Errorf("failed to insert ticket owner data: %w", err)
	}

	return entities.Ticket{
		TicketID:    createdTicket.TicketID,
		BookingID:   createdTicket.BookingID.Int64,
		FlightID:    createdTicket.FlightID,
		Price:       createdTicket.Price,
		FlightClass: entities.FlightClass(createdTicket.FlightClass),
		Owner: entities.TicketOwner{
			FirstName:            ticket.OwnerData.FirstName,
			LastName:             ticket.OwnerData.LastName,
			PhoneNumber:          ticket.OwnerData.PhoneNumber,
			DateOfBirth:          parseDate(ticket.OwnerData.DateOfBirth),
			Gender:               entities.GenderType(ticket.OwnerData.Gender),
			PassportNumber:       ticket.OwnerData.IdentityCardNumber,
			IdentificationNumber: ticket.OwnerData.IdentityCardNumber,
			Address:              ticket.OwnerData.Address,
		},
	}, nil
}

func parseDate(dateStr string) time.Time {
	parsedDate, _ := time.Parse("2006-01-02", dateStr)
	return parsedDate
}
