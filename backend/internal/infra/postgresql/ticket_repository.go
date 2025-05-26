package postgresql

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/spaghetti-lover/qairlines/db/sqlc"
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
		return nil, fmt.Errorf("failed to get tickets by flight ID: %w", err)
	}

	// Chuyển đổi từ DB model sang domain entity
	var result []entities.Ticket
	for _, t := range tickets {
		result = append(result, entities.Ticket{
			TicketID:    strconv.FormatInt(t.FlightID.Int64, 10),
			Status:      entities.TicketStatus(t.Status),
			FlightClass: entities.FlightStatus(t.FlightClass),
			BookingID:   strconv.FormatInt(t.BookingID.Int64, 10),
			FlightID:    strconv.FormatInt(t.FlightID.Int64, 10),
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
