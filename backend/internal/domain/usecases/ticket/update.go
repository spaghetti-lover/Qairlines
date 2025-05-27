package ticket

import (
	"context"
	"errors"
	"strconv"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
)

type IUpdateSeatsUseCase interface {
	Execute(ctx context.Context, updates []dto.UpdateSeatRequest) ([]entities.Ticket, error)
}

type UpdateSeatsUseCase struct {
	ticketRepository adapters.ITicketRepository
}

func NewUpdateSeatsUseCase(ticketRepository adapters.ITicketRepository) IUpdateSeatsUseCase {
	return &UpdateSeatsUseCase{
		ticketRepository: ticketRepository,
	}
}

func (u *UpdateSeatsUseCase) Execute(ctx context.Context, updates []dto.UpdateSeatRequest) ([]entities.Ticket, error) {
	var responses []entities.Ticket

	for _, update := range updates {
		ticketID, err := strconv.ParseInt(update.TicketID, 10, 64)
		if err != nil {
			return nil, err
		}
		ticket, err := u.ticketRepository.UpdateSeat(ctx, ticketID, update.SeatCode)
		if err != nil {
			if errors.Is(err, adapters.ErrTicketNotFound) {
				return nil, adapters.ErrTicketNotFound
			}
			if errors.Is(err, adapters.ErrInvalidSeat) {
				return nil, adapters.ErrInvalidSeat
			}
			return nil, err
		}

		responses = append(responses, *ticket)
	}

	return responses, nil
}
