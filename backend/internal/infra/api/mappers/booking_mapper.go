package mappers

import (
	"strconv"
	"time"

	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
)

func ToCreateBookingParams(request dto.CreateBookingRequest, departureFlight entities.Flight, returnFlight *entities.Flight, email string) entities.CreateBookingParams {
	var returnFlightID string
	if returnFlight != nil {
		returnFlightID = strconv.FormatInt(returnFlight.FlightID, 10)
	}

	return entities.CreateBookingParams{
		Email:                   email,
		DepartureCity:           request.DepartureCity,
		ArrivalCity:             request.ArrivalCity,
		DepartureFlightID:       strconv.FormatInt(departureFlight.FlightID, 10),
		ReturnFlightID:          returnFlightID,
		TripType:                entities.TripType(request.TripType),
		DepartureTicketDataList: mapTicketDataList(request.DepartureTicketDataList),
		ReturnTicketDataList:    mapTicketDataList(request.ReturnTicketDataList),
	}
}

func mapTicketDataList(ticketDataList []dto.TicketDataRequest) []entities.Ticket {
	var mappedList []entities.Ticket
	for _, ticket := range ticketDataList {
		dateOfBirth, err := time.Parse("2006-01-02", ticket.OwnerData.DateOfBirth)
		if err != nil {
			return []entities.Ticket{}
		}
		mappedList = append(mappedList, entities.Ticket{
			Price:       ticket.Price,
			FlightClass: entities.FlightClass(ticket.FlightClass),
			Owner: entities.TicketOwner{
				IdentificationNumber: ticket.OwnerData.IdentityCardNumber,
				FirstName:            ticket.OwnerData.FirstName,
				LastName:             ticket.OwnerData.LastName,
				PhoneNumber:          ticket.OwnerData.PhoneNumber,
				DateOfBirth:          dateOfBirth,
				Gender:               entities.GenderType(ticket.OwnerData.Gender),
				Address:              ticket.OwnerData.Address,
			},
		})
	}
	return mappedList
}

func ToCreateBookingResponse(booking entities.Booking, departureTickets []entities.Ticket, returnTickets []entities.Ticket) dto.CreateBookingResponse {
	var returnTicketsResponse []dto.TicketDataResponse
	if returnTickets != nil {
		returnTicketsResponse = mapTicketDataListToResponse(returnTickets)
	}

	return dto.CreateBookingResponse{
		BookingID:         strconv.FormatInt(booking.BookingID, 10),
		DepartureFlightID: strconv.FormatInt(booking.DepartureFlightID, 10),
		ReturnFlightID:    mapNullableInt64ToString(booking.ReturnFlightID),
		TripType:          string(booking.TripType),
		DepartureTickets:  mapTicketDataListToResponse(departureTickets),
		ReturnTickets:     returnTicketsResponse,
	}
}

func mapTicketDataListToResponse(ticketDataList []entities.Ticket) []dto.TicketDataResponse {
	var mappedList []dto.TicketDataResponse
	for _, ticket := range ticketDataList {
		seatID := ""
		if ticket.SeatID != 0 {
			seatID = strconv.FormatInt(ticket.SeatID, 10)
		}
		mappedList = append(mappedList, dto.TicketDataResponse{
			TicketID:    strconv.FormatInt(ticket.TicketID, 10),
			SeatID:      seatID,
			Price:       ticket.Price,
			FlightClass: string(ticket.FlightClass),
			OwnerData: dto.OwnerData{
				IdentityCardNumber: ticket.Owner.IdentificationNumber,
				FirstName:          ticket.Owner.FirstName,
				LastName:           ticket.Owner.LastName,
				PhoneNumber:        ticket.Owner.PhoneNumber,
				DateOfBirth:        ticket.Owner.DateOfBirth.String(),
				Gender:             ticket.Owner.Gender,
				Address:            ticket.Owner.Address,
			},
		})
	}
	return mappedList
}

func mapNullableInt64ToString(value *int64) string {
	if value == nil {
		return ""
	}
	return strconv.FormatInt(*value, 10)
}

func ToGetBookingResponse(booking entities.Booking, departureTickets []entities.Ticket, returnTickets []entities.Ticket) dto.GetBookingResponse {
	return dto.GetBookingResponse{
		BookingID:         strconv.FormatInt(booking.BookingID, 10),
		Email:             booking.UserEmail,
		TripType:          string(booking.TripType),
		DepartureFlightID: strconv.FormatInt(booking.DepartureFlightID, 10),
		ReturnFlightID:    mapNullableInt64ToString(booking.ReturnFlightID),
		DepartureTickets:  mapTicketIDsToResponse(departureTickets),
		ReturnTickets:     mapTicketIDsToResponse(returnTickets),
		CreatedAt:         booking.CreatedAt.Format(time.RFC3339),
		UpdatedAt:         booking.UpdatedAt.Format(time.RFC3339),
	}
}

func mapTicketIDsToResponse(tickets []entities.Ticket) []string {
	var ticketIDs []string
	for _, ticket := range tickets {
		ticketIDs = append(ticketIDs, strconv.FormatInt(ticket.TicketID, 10))
	}
	return ticketIDs
}
