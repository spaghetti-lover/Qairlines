package postgresql

import (
	"context"
	"database/sql"

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

func (r *TicketRepositoryPostgres) GetTicketByID(ctx context.Context, ticketID int64) (*entities.Ticket, error) {
	// Sử dụng sqlc để lấy vé theo ticketID
	ticket, err := r.store.GetTicketByID(ctx, ticketID)
	if err == sql.ErrNoRows {
		return nil, adapters.ErrTicketNotFound
	}
	if err != nil {
		return nil, err
	}

	return &entities.Ticket{
		TicketID:    ticket.TicketID,
		Status:      entities.TicketStatus(ticket.Status),
		FlightClass: entities.FlightClass(ticket.FlightClass),
		Price:       int(ticket.Price),
		BookingID:   ticket.BookingID.Int64,
		FlightID:    ticket.FlightID,
		CreatedAt:   ticket.CreatedAt,
		UpdatedAt:   ticket.UpdatedAt,
		Seat: entities.Seat{
			SeatCode: ticket.SeatCode.String,
		},
		Owner: entities.TicketOwner{
			FirstName:   ticket.OwnerFirstName.String,
			LastName:    ticket.OwnerLastName.String,
			PhoneNumber: ticket.OwnerPhoneNumber.String,
			Gender:      entities.GenderType(ticket.OwnerGender.GenderType),
		},
	}, nil
}

func (r *TicketRepositoryPostgres) CancelTicket(ctx context.Context, ticketID int64) (*entities.Ticket, error) {
	row, err := r.store.CancelTicket(ctx, ticketID)
	if err == sql.ErrNoRows {
		return nil, adapters.ErrTicketNotFound
	}
	if err != nil {
		return nil, adapters.ErrTicketCannotBeCancelled
	}

	return &entities.Ticket{
		TicketID:    row.TicketID,
		Status:      entities.TicketStatus(row.Status),
		FlightClass: entities.FlightClass(row.FlightClass),
		Price:       int(row.Price),
		BookingID:   row.BookingID.Int64,
		FlightID:    row.FlightID,
		UpdatedAt:   row.UpdatedAt,
		Seat: entities.Seat{
			SeatCode: row.SeatCode,
		},
		Owner: entities.TicketOwner{
			FirstName:   row.OwnerFirstName.String,
			LastName:    row.OwnerLastName.String,
			PhoneNumber: row.OwnerPhoneNumber.String,
		},
	}, nil
}
