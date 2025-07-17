package booking

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/hibiken/asynq"
	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
	"github.com/spaghetti-lover/qairlines/internal/infra/worker"
)

type ICreateBookingUseCase interface {
	Execute(ctx context.Context, booking dto.CreateBookingRequest, email string) (dto.CreateBookingResponse, error)
}

type CreateBookingUseCase struct {
	bookingRepository adapters.IBookingRepository
	flightRepository  adapters.IFlightRepository
	taskDistributor   worker.TaskDistributor
}

func NewCreateBookingUseCase(bookingRepository adapters.IBookingRepository, flightRepository adapters.IFlightRepository, taskDistributor worker.TaskDistributor) ICreateBookingUseCase {
	return &CreateBookingUseCase{
		bookingRepository: bookingRepository,
		flightRepository:  flightRepository,
		taskDistributor:   taskDistributor,
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
			return dto.CreateBookingResponse{}, err
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
	params := mappers.ToCreateBookingParams(booking, *departureFlight, returnFlight, email)
	arg := entities.CreateBookingParams{
		Email:                   params.Email,
		DepartureCity:           params.DepartureCity,
		ArrivalCity:             params.ArrivalCity,
		DepartureFlightID:       params.DepartureFlightID,
		ReturnFlightID:          params.ReturnFlightID,
		TripType:                params.TripType,
		DepartureTicketDataList: params.DepartureTicketDataList,
		ReturnTicketDataList:    params.ReturnTicketDataList,
		AfterCreate: func(booking entities.Booking) error {
			taskPayload := &worker.PayloadSendVerifyEmail{
				To:      booking.UserEmail,
				Subject: "Xác nhận ghế máy bay",
				Body: fmt.Sprintf(
					`<html>
						<body>
							<h2>Xin chào,</h2>
							<p>Chúng tôi xin thông báo rằng ghế của bạn đã được <b>cập nhật thành công</b> cho chuyến bay.</p>
							<p><strong>Mã chuyến bay:</strong> %d</p>
							<p>Vui lòng kiểm tra lại thông tin trong ứng dụng để đảm bảo mọi thứ chính xác.</p>
							<p>Chúc bạn có một chuyến bay an toàn và thoải mái!</p>
							<br>
							<p>Trân trọng,<br>
							<b>Đội ngũ Qairlines</b></p>
						</body>
						</html>`,
					booking.BookingID,
				),
			}
			opts := []asynq.Option{
				asynq.MaxRetry(10),
				asynq.ProcessIn(10 * time.Second),
				asynq.Queue(worker.QueueCritical),
			}
			return u.taskDistributor.DistributeTaskSendVerifyEmail(ctx, taskPayload, opts...)
		},
	}
	createdBooking, departureTickets, returnTickets, err := u.bookingRepository.CreateBookingTx(ctx, arg)

	if err != nil {
		return dto.CreateBookingResponse{}, err
	}

	// Map kết quả sang DTO
	return mappers.ToCreateBookingResponse(createdBooking, departureTickets, returnTickets), nil
}
