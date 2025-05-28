package postgresql

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/spaghetti-lover/qairlines/db/sqlc"
	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type BookingRepositoryPostgres struct {
	store db.Store
}

func NewBookingRepositoryPostgres(store *db.Store) *BookingRepositoryPostgres {
	return &BookingRepositoryPostgres{
		store: *store,
	}
}

func (r *BookingRepositoryPostgres) CreateBookingTx(ctx context.Context, booking entities.CreateBookingParams) (entities.Booking, []entities.Ticket, []entities.Ticket, error) {
	// Chuyển đổi từ entities.CreateBookingParams sang db.CreateBookingTxParams
	departureID, err := strconv.ParseInt(booking.DepartureFlightID, 10, 64)
	if err != nil {
		return entities.Booking{}, nil, nil, err
	}
	returnFlight, err := strconv.ParseInt(booking.ReturnFlightID, 10, 64)
	if err != nil && booking.TripType == entities.RoundTrip {
		return entities.Booking{}, nil, nil, adapters.ErrFlightNotFound
	} else if err != nil {
		return entities.Booking{}, nil, nil, adapters.ErrFlightNotFound
	} else if booking.TripType == entities.RoundTrip && returnFlight == 0 {
		return entities.Booking{}, nil, nil, adapters.ErrFlightNotFound
	}
	txParams := db.CreateBookingTxParams{
		UserEmail:           booking.Email,
		TripType:            string(booking.TripType),
		DepartureFlightID:   departureID,
		ReturnFlightID:      &returnFlight,
		DepartureTicketData: booking.DepartureTicketDataList,
		ReturnTicketData:    booking.ReturnTicketDataList,
	}

	// Gọi CreateBookingTx từ tầng SQLStore
	txResult, err := r.store.CreateBookingTx(ctx, txParams)
	if err != nil {
		return entities.Booking{}, nil, nil, err
	}

	// Trả về kết quả
	return txResult.Booking, txResult.DepartureTickets, txResult.ReturnTickets, nil
}

func (r *BookingRepositoryPostgres) GetBookingByID(ctx context.Context, bookingID int64) (entities.Booking, []entities.Ticket, []entities.Ticket, error) {
	// Lấy thông tin booking
	booking, err := r.store.GetBooking(ctx, bookingID)
	if err != nil {
		return entities.Booking{}, nil, nil, err
	}

	// Lấy danh sách vé cho chuyến bay đi
	departureTickets, err := r.store.GetTicketsByBookingIDAndType(ctx, db.GetTicketsByBookingIDAndTypeParams{
		BookingID: pgtype.Int8{Int64: booking.BookingID, Valid: true},
		Column2:   "departure",
	})
	if err != nil {
		return entities.Booking{}, nil, nil, err
	}

	// Lấy danh sách vé cho chuyến bay về
	returnTickets, err := r.store.GetTicketsByBookingIDAndType(ctx, db.GetTicketsByBookingIDAndTypeParams{
		BookingID: pgtype.Int8{Int64: booking.BookingID, Valid: true},
		Column2:   "return",
	})
	if err != nil {
		return entities.Booking{}, nil, nil, err
	}

	return entities.Booking{
		BookingID:         booking.BookingID,
		UserEmail:         booking.UserEmail.String,
		TripType:          entities.TripType(booking.TripType),
		DepartureFlightID: booking.DepartureFlightID.Int64,
		ReturnFlightID:    &booking.ReturnFlightID.Int64,
		CreatedAt:         booking.CreatedAt,
		UpdatedAt:         booking.UpdatedAt,
		Status:            entities.BookingStatus(booking.Status),
	}, mapDBTicketsToEntitiesTickets(departureTickets), mapDBTicketsToEntitiesTickets(returnTickets), nil
}

func mapDBTicketsToEntitiesTickets(dbTickets []db.Ticket) []entities.Ticket {
	var entityTickets []entities.Ticket
	for _, dbTicket := range dbTickets {
		entityTickets = append(entityTickets, entities.Ticket{
			TicketID:    dbTicket.TicketID,
			SeatID:      sql.NullInt64{Int64: dbTicket.SeatID.Int64, Valid: dbTicket.SeatID.Valid},
			FlightClass: entities.FlightClass(dbTicket.FlightClass),
			Price:       dbTicket.Price,
			Status:      entities.TicketStatus(dbTicket.Status),
			BookingID:   dbTicket.BookingID.Int64,
			FlightID:    dbTicket.FlightID,
			CreatedAt:   dbTicket.CreatedAt,
			UpdatedAt:   dbTicket.UpdatedAt,
		})
	}
	return entityTickets
}
