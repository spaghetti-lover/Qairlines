package db

import (
	"context"
	"fmt"
)

type SeatUpdateParams struct {
	TicketID int64
	SeatCode string
}

// BookingTx performs a booking transfer from user
func (store *SQLStore) UpdateSeats(ctx context.Context, bookingID int64, seats []SeatUpdateParams) error {
	err := store.execTx(ctx, func(q *Queries) error {
		for _, seat := range seats {
			// Kiểm tra ghế có thuộc chuyến bay không
			ticket, err := store.GetTicketByID(ctx, seat.TicketID)
			if err != nil || ticket.BookingID == nil || *ticket.BookingID != bookingID {
				return fmt.Errorf("invalid ticket_id %d for booking_id %d", seat.TicketID, bookingID)
			}

			// Kiểm tra ghế có còn trống không
			isAvailable, err := store.CheckSeatAvailability(ctx, CheckSeatAvailabilityParams{
				SeatCode: seat.SeatCode,
				FlightID: &ticket.FlightID,
			})
			if err != nil || !isAvailable {
				return fmt.Errorf("seat %s is not available", seat.SeatCode)
			}

			// Đánh dấu ghế là không còn trống trong bảng Seats
			err = q.MarkSeatUnavailable(ctx, MarkSeatUnavailableParams{
				FlightID: &ticket.FlightID,
				SeatCode: seat.SeatCode,
			})
			if err != nil {
				return fmt.Errorf("failed to mark seat %s as unavailable: %w", seat.SeatCode, err)
			}
		}
		return nil
	})
	return err
}
