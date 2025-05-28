package postgresql

import (
	"context"
	"strconv"

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
