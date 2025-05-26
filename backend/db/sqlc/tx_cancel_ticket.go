package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// CancelTicketTxParams chứa thông tin cần thiết để hủy vé
type CancelTicketTxParams struct {
	TicketID int64 `json:"ticket_id"`
}

// CancelTicketTxResult chứa kết quả của transaction hủy vé
type CancelTicketTxResult struct {
	Ticket          GetTicketByIDRow `json:"ticket"`
	UpdatedSeatID   int64            `json:"updated_seat_id"`
	Success         bool             `json:"success"`
	TransactionTime time.Time        `json:"transaction_time"`
}

// CancelTicketTx thực hiện transaction hủy vé và cập nhật trạng thái ghế
func (store *SQLStore) CancelTicketTx(ctx context.Context, arg CancelTicketTxParams) (CancelTicketTxResult, error) {
	var result CancelTicketTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// 1. Kiểm tra xem vé có tồn tại không
		ticket, err := q.GetTicketByID(ctx, arg.TicketID)
		if err != nil {
			if err == sql.ErrNoRows {
				return fmt.Errorf("ticket with ID %d not found", arg.TicketID)
			}
			return err
		}

		// 2. Kiểm tra xem vé có thể hủy được không
		if ticket.Status != "booked" {
			return fmt.Errorf("ticket with ID %d cannot be cancelled due to its current status: %s", arg.TicketID, ticket.Status)
		}

		// 3. Cập nhật trạng thái vé
		_, err = q.UpdateTicketStatus(ctx, UpdateTicketStatusParams{
			TicketID: arg.TicketID,
			Status:   "cancelled",
		})
		if err != nil {
			return fmt.Errorf("failed to update ticket status: %w", err)
		}

		// 4. Cập nhật trạng thái ghế thành có sẵn
		err = q.UpdateSeatAvailability(ctx, UpdateSeatAvailabilityParams{
			TicketID:    ticket.TicketID,
			IsAvailable: true,
		})
		if err != nil {
			return fmt.Errorf("failed to update seat availability: %w", err)
		}

		// 5. Lấy thông tin vé đã cập nhật đầy đủ cho kết quả
		updatedTicketDetails, err := q.GetTicketByID(ctx, arg.TicketID)
		if err != nil {
			return fmt.Errorf("failed to get updated ticket details: %w", err)
		}

		// 6. Cập nhật kết quả
		result.Ticket = updatedTicketDetails
		result.UpdatedSeatID = ticket.SeatID.Int64
		result.Success = true
		result.TransactionTime = time.Now()

		return nil
	})

	return result, err
}
