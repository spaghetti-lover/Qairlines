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

	departureTicketData := make([]db.TicketData, len(booking.DepartureTicketDataList))
	for i, ticketData := range booking.DepartureTicketDataList {
		departureTicketData[i] = db.TicketData{
			Price:       int64(ticketData.Price),
			FlightClass: string(ticketData.FlightClass),
			OwnerData: db.OwnerData{
				IdentityCardNumber: ticketData.Owner.IdentificationNumber,
				FirstName:          ticketData.Owner.FirstName,
				LastName:           ticketData.Owner.LastName,
				PhoneNumber:        ticketData.Owner.PhoneNumber,
				DateOfBirth:        ticketData.Owner.DateOfBirth.String(),
				Gender:             string(ticketData.Owner.Gender),
				Address:            ticketData.Owner.Address,
			},
		}
	}
	returnTicketDataList := make([]db.TicketData, len(booking.ReturnTicketDataList))
	var returnFlightID int64
	if booking.ReturnFlightID != "" {
		returnFlightID, err = strconv.ParseInt(booking.ReturnFlightID, 10, 64)

		if err != nil && booking.TripType == entities.RoundTrip {
			return entities.Booking{}, nil, nil, adapters.ErrFlightNotFound
		} else if err != nil {
			return entities.Booking{}, nil, nil, adapters.ErrFlightNotFound
		} else if booking.TripType == entities.RoundTrip && returnFlightID == 0 {
			return entities.Booking{}, nil, nil, adapters.ErrFlightNotFound
		}
		for i, ticketData := range booking.ReturnTicketDataList {
			returnTicketDataList[i] = db.TicketData{
				Price:       int64(ticketData.Price),
				FlightClass: string(ticketData.FlightClass),
				OwnerData: db.OwnerData{
					IdentityCardNumber: ticketData.Owner.IdentificationNumber,
					FirstName:          ticketData.Owner.FirstName,
					LastName:           ticketData.Owner.LastName,
					PhoneNumber:        ticketData.Owner.PhoneNumber,
					DateOfBirth:        ticketData.Owner.DateOfBirth.String(),
					Gender:             string(ticketData.Owner.Gender),
					Address:            ticketData.Owner.Address,
				},
			}
		}
	}

	txParams := db.CreateBookingTxParams{
		UserEmail:           booking.Email,
		TripType:            string(booking.TripType),
		DepartureFlightID:   departureID,
		ReturnFlightID:      returnFlightID,
		DepartureTicketData: departureTicketData,
		ReturnTicketData:    returnTicketDataList,
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
		BookingID: &booking.BookingID,
		Column2:   "departure",
	})
	if err != nil {
		return entities.Booking{}, nil, nil, err
	}

	// Lấy danh sách vé cho chuyến bay về
	returnTickets, err := r.store.GetTicketsByBookingIDAndType(ctx, db.GetTicketsByBookingIDAndTypeParams{
		BookingID: &booking.BookingID,
		Column2:   "return",
	})
	if err != nil {
		return entities.Booking{}, nil, nil, err
	}

	return entities.Booking{
		BookingID:         booking.BookingID,
		UserEmail:         *booking.UserEmail,
		TripType:          entities.TripType(booking.TripType),
		DepartureFlightID: *booking.DepartureFlightID,
		ReturnFlightID:    booking.ReturnFlightID,
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
			SeatID:      dbTicket.SeatID,
			FlightClass: entities.FlightClass(dbTicket.FlightClass),
			Price:       dbTicket.Price,
			Status:      entities.TicketStatus(dbTicket.Status),
			BookingID:   *dbTicket.BookingID,
			FlightID:    dbTicket.FlightID,
			CreatedAt:   dbTicket.CreatedAt,
			UpdatedAt:   dbTicket.UpdatedAt,
		})
	}
	return entityTickets
}
