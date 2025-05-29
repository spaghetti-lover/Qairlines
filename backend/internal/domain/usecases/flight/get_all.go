package flight

import (
	"context"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
)

type IGetAllFlightsUseCase interface {
	Execute(ctx context.Context) ([]entities.Flight, []entities.Ticket, error)
}

type GetAllFlightsUseCase struct {
	flightRepository adapters.IFlightRepository
	ticketRepository adapters.ITicketRepository
}

func NewGetAllFlightsUseCase(flightRepository adapters.IFlightRepository, ticketRepository adapters.ITicketRepository) IGetAllFlightsUseCase {
	return &GetAllFlightsUseCase{
		flightRepository: flightRepository,
		ticketRepository: ticketRepository,
	}
}

func (u *GetAllFlightsUseCase) Execute(ctx context.Context) ([]entities.Flight, []entities.Ticket, error) {
	// Lấy danh sách chuyến bay từ repository
	flights, err := u.flightRepository.GetAllFlights(ctx)
	tickets := []entities.Ticket{}
	for _, flight := range flights {
		ticketList, err := u.ticketRepository.GetTicketsByFlightID(ctx, flight.FlightID)
		if err != nil {
			return nil, nil, err
		}
		tickets = append(tickets, ticketList...)
	}
	if err != nil {
		return nil, nil, err
	}

	// Sử dụng mapper để chuyển đổi danh sách entity sang DTO
	return flights, tickets, nil
}
