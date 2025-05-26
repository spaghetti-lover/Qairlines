package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/spaghetti-lover/qairlines/db/sqlc"
	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type TicketRepositoryPostgres struct {
	store db.Store
}

// NewTicketRepositoryPostgres creates a new TicketRepositoryPostgres instance
func NewTicketRepositoryPostgres(store *db.Store) *TicketRepositoryPostgres {
	return &TicketRepositoryPostgres{
		store: *store,
	}
}

func (r *TicketRepositoryPostgres) GetTicketsByFlightID(ctx context.Context, flightID int64) ([]entities.Ticket, error) {
	// Sử dụng query đã được tạo trong SQLC
	tickets, err := r.store.GetTicketsByFlightID(ctx, pgtype.Int8{Int64: flightID, Valid: true})
	if err != nil {
		return nil, adapters.ErrFlightNotFound
	}

	// Chuyển đổi từ DB model sang domain entity
	var result []entities.Ticket
	for _, t := range tickets {
		result = append(result, entities.Ticket{
			TicketID:    strconv.FormatInt(t.FlightID.Int64, 10),
			Status:      entities.TicketStatus(t.Status),
			FlightClass: entities.FlightClass(t.FlightClass),
			BookingID:   t.BookingID.Int64,
			FlightID:    t.FlightID.Int64,
			CreatedAt:   t.CreatedAt,
			UpdatedAt:   t.UpdatedAt,
			Owner: entities.TicketOwner{
				FirstName:   t.FirstName.String,
				LastName:    t.LastName.String,
				PhoneNumber: t.PhoneNumber.String,
				Gender:      string(t.Gender.GenderType),
			},
			Seat: entities.Seat{
				SeatCode: t.SeatCode.String,
				Class:    string(t.SeatClass.FlightClass),
			},
		})
	}

	return result, nil
}

func (r *TicketRepositoryPostgres) GetTicketByID(ctx context.Context, ticketID int64) (*entities.Ticket, error) {
	ticket, err := r.store.GetTicketByID(ctx, ticketID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, adapters.ErrTicketNotFound
		}
		return nil, fmt.Errorf("failed to get ticket by ID: %w", err)
	}

	return &entities.Ticket{
		TicketID:    strconv.FormatInt(ticket.TicketID, 10),
		Status:      entities.TicketStatus(ticket.Status),
		SeatCode:    ticket.SeatCode,
		FlightClass: entities.FlightClass(ticket.FlightClass),
		Price:       ticket.Price,
		BookingID:   ticket.BookingID.Int64,
		FlightID:    ticket.FlightID.Int64,
		CreatedAt:   ticket.CreatedAt,
		UpdatedAt:   ticket.UpdatedAt,
		Owner: entities.TicketOwner{
			FirstName:   ticket.FirstName.String,
			LastName:    ticket.LastName.String,
			PhoneNumber: ticket.PhoneNumber.String,
			Gender:      string(ticket.Gender),
		},
	}, nil
}

func (r *TicketRepositoryPostgres) CancelTicket(ctx context.Context, ticketID int64) (*entities.Ticket, error) {
	// Sử dụng transaction một lần duy nhất
	result, err := r.store.CancelTicketTx(ctx, db.CancelTicketTxParams{
		TicketID: ticketID,
	})

	if err != nil {
		if err.Error() == fmt.Sprintf("ticket with ID %d not found", ticketID) {
			return nil, adapters.ErrTicketNotFound
		}
		if err.Error() == fmt.Sprintf("ticket with ID %d cannot be cancelled due to its current status: %s", ticketID, "cancelled") {
			return nil, adapters.ErrTicketCannotBeCancelled
		}
		return nil, fmt.Errorf("failed to cancel ticket: %w", err)
	}

	// Chuyển đổi dữ liệu từ kết quả transaction thành entity
	ticket := result.Ticket
	return &entities.Ticket{
		TicketID:    strconv.FormatInt(ticket.TicketID, 10),
		Status:      entities.TicketStatus(ticket.Status),
		SeatCode:    ticket.SeatCode,
		FlightClass: entities.FlightClass(ticket.FlightClass),
		Price:       ticket.Price,
		BookingID:   ticket.BookingID.Int64,
		FlightID:    ticket.FlightID.Int64,
		CreatedAt:   ticket.CreatedAt,
		UpdatedAt:   ticket.UpdatedAt,
		Owner: entities.TicketOwner{
			FirstName:   ticket.FirstName.String,
			LastName:    ticket.LastName.String,
			PhoneNumber: ticket.PhoneNumber.String,
			Gender:      string(ticket.Gender),
		},
	}, nil
}
