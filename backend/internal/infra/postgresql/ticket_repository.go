package postgresql

import (
	"context"

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
	// Sử dụng sqlc để lấy danh sách vé theo flightID
	tickets, err := r.store.GetTicketsByFlightID(ctx, flightID)
	if err != nil {
		return nil, err
	}

	// Map từ kết quả sqlc sang entity Ticket
	var result []entities.Ticket
	for _, t := range tickets {
		result = append(result, entities.Ticket{
			TicketID:    t.TicketID,
			SeatID:      t.SeatID,
			FlightClass: entities.FlightClass(t.FlightClass),
			Price:       int(t.Price),
			Status:      entities.TicketStatus(t.Status),
			BookingID:   t.BookingID.Int64,
			FlightID:    t.FlightID,
			CreatedAt:   t.CreatedAt,
			UpdatedAt:   t.UpdatedAt,
			Seat: entities.Seat{
				SeatID:      t.SeatID,
				SeatCode:    t.SeatCode.String,
				IsAvailable: t.IsAvailable.Bool,
				Class:       entities.FlightClass(t.SeatClass.FlightClass),
			},
			Owner: entities.TicketOwner{
				FirstName:   t.OwnerFirstName.String,
				LastName:    t.OwnerLastName.String,
				PhoneNumber: t.OwnerPhoneNumber.String,
				Gender:      entities.GenderType(t.OwnerGender.GenderType),
			},
		})
	}

	return result, nil
}
