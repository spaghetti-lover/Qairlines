package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type CreateBookingTxParams struct {
	UserEmail           string
	TripType            string
	DepartureFlightID   int64
	ReturnFlightID      *int64
	DepartureTicketData []entities.Ticket
	ReturnTicketData    []entities.Ticket
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
		booking, err := q.CreateBooking(ctx, CreateBookingParams{
			UserEmail:         pgtype.Text{String: arg.UserEmail, Valid: true},
			TripType:          TripType(arg.TripType),
			DepartureFlightID: pgtype.Int8{Int64: arg.DepartureFlightID, Valid: true},
			ReturnFlightID:    pgtype.Int8{Int64: *arg.ReturnFlightID, Valid: arg.ReturnFlightID != nil},
			Status:            BookingStatus(entities.BookingStatusPending),
		})
		if err != nil {
			return fmt.Errorf("failed to create booking: %w", err)
		}
		result.Booking = entities.Booking{
			BookingID:         booking.BookingID,
			UserEmail:         booking.UserEmail.String,
			TripType:          entities.TripType(booking.TripType),
			DepartureFlightID: booking.DepartureFlightID.Int64,
			ReturnFlightID:    &booking.ReturnFlightID.Int64,
			Status:            entities.BookingStatus(booking.Status),
			CreatedAt:         booking.CreatedAt,
			UpdatedAt:         booking.UpdatedAt,
		}

		// Tạo vé cho chuyến bay đi
		for _, ticket := range arg.DepartureTicketData {
			createdTicket, err := q.CreateTicket(ctx, CreateTicketParams{
				FlightClass: FlightClass(ticket.FlightClass),
				Price:       ticket.Price,
				Status:      TicketStatusBooked,
				BookingID:   pgtype.Int8{Int64: booking.BookingID, Valid: true},
				FlightID:    booking.DepartureFlightID.Int64,
			})
			if err != nil {
				return fmt.Errorf("failed to create departure ticket: %w", err)
			}
			// Lưu thông tin chủ sở hữu vé vào TicketOwnerSnapshot
			_, err = q.CreateTicketOwnerSnapshot(ctx, CreateTicketOwnerSnapshotParams{
				TicketID:             createdTicket.TicketID,
				FirstName:            pgtype.Text{String: ticket.Owner.FirstName, Valid: true},
				LastName:             pgtype.Text{String: ticket.Owner.LastName, Valid: true},
				PhoneNumber:          pgtype.Text{String: ticket.Owner.PhoneNumber, Valid: true},
				Gender:               GenderType(ticket.Owner.Gender),
				DateOfBirth:          pgtype.Date{Time: ticket.Owner.DateOfBirth, Valid: true},
				PassportNumber:       pgtype.Text{String: ticket.Owner.PassportNumber, Valid: true},
				IdentificationNumber: pgtype.Text{String: ticket.Owner.IdentificationNumber, Valid: true},
				Address:              pgtype.Text{String: ticket.Owner.Address, Valid: true},
			})
			if err != nil {
				return fmt.Errorf("failed to insert ticket owner data: %w", err)
			}

			result.DepartureTickets = append(result.DepartureTickets, entities.Ticket{
				TicketID:    createdTicket.TicketID,
				BookingID:   createdTicket.BookingID.Int64,
				FlightID:    createdTicket.FlightID,
				Price:       createdTicket.Price,
				FlightClass: entities.FlightClass(createdTicket.FlightClass),
				Owner:       ticket.Owner,
			})
		}

		// Tạo vé cho chuyến bay về (nếu có)
		if arg.TripType == string(entities.RoundTrip) && arg.ReturnFlightID != nil {
			for _, ticket := range arg.ReturnTicketData {
				createdTicket, err := q.CreateTicket(ctx, CreateTicketParams{
					FlightClass: FlightClass(ticket.FlightClass),
					Price:       ticket.Price,
					Status:      TicketStatusBooked,
					BookingID:   pgtype.Int8{Int64: booking.BookingID, Valid: true},
					FlightID:    *arg.ReturnFlightID,
				})
				if err != nil {
					return fmt.Errorf("failed to create return ticket: %w", err)
				}

				_, err = q.db.Exec(ctx, `
                    INSERT INTO TicketOwnerSnapshot (
                        ticket_id, first_name, last_name, phone_number, gender, date_of_birth,
                        passport_number, identification_number, address
                    ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
                `, createdTicket.TicketID,
					ticket.Owner.FirstName,
					ticket.Owner.LastName,
					ticket.Owner.PhoneNumber,
					string(ticket.Owner.Gender),
					ticket.Owner.DateOfBirth,
					ticket.Owner.PassportNumber,
					ticket.Owner.IdentificationNumber,
					ticket.Owner.Address,
				)
				if err != nil {
					return fmt.Errorf("failed to insert ticket owner data: %w", err)
				}

				result.ReturnTickets = append(result.ReturnTickets, entities.Ticket{
					TicketID:    createdTicket.TicketID,
					BookingID:   createdTicket.BookingID.Int64,
					FlightID:    createdTicket.FlightID,
					Price:       createdTicket.Price,
					FlightClass: entities.FlightClass(createdTicket.FlightClass),
					Owner:       ticket.Owner,
				})
			}
		}

		return nil
	})

	return result, err
}
