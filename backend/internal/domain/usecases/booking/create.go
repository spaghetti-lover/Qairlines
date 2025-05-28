package booking

import (
	"context"
	"errors"
	"strconv"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
)

type ICreateBookingUseCase interface {
	Execute(ctx context.Context, booking dto.CreateBookingRequest, email string) (dto.CreateBookingResponse, error)
}

type CreateBookingUseCase struct {
	bookingRepository adapters.IBookingRepository
	flightRepository  adapters.IFlightRepository
}

func NewCreateBookingUseCase(bookingRepository adapters.IBookingRepository, flightRepository adapters.IFlightRepository) ICreateBookingUseCase {
	return &CreateBookingUseCase{
		bookingRepository: bookingRepository,
		flightRepository:  flightRepository,
	}
}

func (u *CreateBookingUseCase) Execute(ctx context.Context, booking dto.CreateBookingRequest, email string) (dto.CreateBookingResponse, error) {
	// Kiểm tra sự tồn tại của chuyến bay đi
	departureFlightID, err := strconv.ParseInt(booking.DepartureFlightID, 10, 64)
	if err != nil {
		return dto.CreateBookingResponse{}, adapters.ErrFlightNotFound
	}
	departureFlight, err := u.flightRepository.GetFlightByID(ctx, departureFlightID)
	if err != nil {
		if errors.Is(err, adapters.ErrFlightNotFound) {
			return dto.CreateBookingResponse{}, adapters.ErrFlightNotFound
		}
		return dto.CreateBookingResponse{}, err
	}

	// Kiểm tra sự tồn tại của chuyến bay về (nếu là roundTrip)
	var returnFlight *entities.Flight
	if booking.TripType == "roundTrip" {
		returnFlightID, err := strconv.ParseInt(booking.ReturnFlightID, 10, 64)
		if err != nil {
			return dto.CreateBookingResponse{}, adapters.ErrFlightNotFound
		}
		returnFlight, err = u.flightRepository.GetFlightByID(ctx, returnFlightID)
		if err != nil {
			if errors.Is(err, adapters.ErrFlightNotFound) {
				return dto.CreateBookingResponse{}, adapters.ErrFlightNotFound
			}
			return dto.CreateBookingResponse{}, err
		}
	}

	// Tạo booking trong repository
	createdBooking, departureTickets, returnTickets, err := u.bookingRepository.CreateBookingTx(ctx, mappers.ToCreateBookingParams(booking, *departureFlight, returnFlight, email))

	if err != nil {
		return dto.CreateBookingResponse{}, err
	}

	// Map kết quả sang DTO
	return mappers.ToCreateBookingResponse(createdBooking, departureTickets, returnTickets), nil
}
