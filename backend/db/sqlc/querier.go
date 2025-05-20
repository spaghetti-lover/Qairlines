package db

import "context"

type Querier interface {
	// Airport
	CreateAirport(ctx context.Context, arg CreateAirportParams) (Airport, error)
	DeleteAirport(ctx context.Context, airportCode string) error
	GetAirport(ctx context.Context, airportCode string) (Airport, error)
	ListAirports(ctx context.Context, arg ListAirportsParams) ([]Airport, error)

	// Airplane
	CreateAirplane(ctx context.Context, arg CreateAirplaneParams) (Airplane, error)
	DeleteAirplane(ctx context.Context, registrationNumber string) error
	GetAirplane(ctx context.Context, registrationNumber string) (Airplane, error)
	ListAirplanes(ctx context.Context, arg ListAirplanesParams) ([]Airplane, error)

	// Booking
	CreateBooking(ctx context.Context, arg CreateBookingParams) (Booking, error)
	DeleteBookings(ctx context.Context, bookingID int64) error
	GetBooking(ctx context.Context, bookingID int64) (Booking, error)
	ListBookings(ctx context.Context, arg ListBookingsParams) ([]Booking, error)

	// Flight Seat
	CreateFlightSeat(ctx context.Context, arg CreateFlightSeatParams) (FlightSeat, error)
	DeleteFlightSeat(ctx context.Context, flightID int64) error
	GetFlightSeat(ctx context.Context, arg GetFlightSeatParams) (FlightSeat, error)
	ListFlightSeats(ctx context.Context, arg ListFlightSeatsParams) ([]FlightSeat, error)
	ListFlightSeatsByFlightID(ctx context.Context, flightID int64) ([]FlightSeat, error)

	// Flight
	CreateFlight(ctx context.Context, arg CreateFlightParams) (Flight, error)
	DeleteFlight(ctx context.Context, flightNumber string) error
	GetFlight(ctx context.Context, flightID int64) (Flight, error)
	ListFlights(ctx context.Context, arg ListFlightsParams) ([]Flight, error)

	// Check seat
	CheckSeatOccupied(ctx context.Context, arg CheckSeatOccupiedParams) (bool, error)
	CountOccupiedSeats(ctx context.Context, arg CountOccupiedSeatsParams) (int64, error)

	// Passenger
	CreatePassenger(ctx context.Context, arg CreatePassengerParams) (Passenger, error)
	DeletePassenger(ctx context.Context, passengerID int64) error
	GetPassenger(ctx context.Context, passengerID int64) (Passenger, error)
	ListPassengers(ctx context.Context, arg ListPassengersParams) ([]Passenger, error)

	// Payment
	CreatePayment(ctx context.Context, arg CreatePaymentParams) (Payment, error)
	DeletePayment(ctx context.Context, paymentID int64) error
	GetPayment(ctx context.Context, paymentID int64) (Payment, error)
	ListPayment(ctx context.Context, arg ListPaymentParams) ([]Payment, error)

	// Airplane Model
	CreateAirplaneModel(ctx context.Context, arg CreateAirplaneModelParams) (AirplaneModel, error)
	DeleteAirplaneModel(ctx context.Context, airplaneModelID int64) error
	GetAirplaneModel(ctx context.Context, airplaneModelID int64) (AirplaneModel, error)
	ListAirplaneModels(ctx context.Context, arg ListAirplaneModelsParams) ([]AirplaneModel, error)

	// User
	GetUser(ctx context.Context, userID int64) (User, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteUser(ctx context.Context, userID int64) error
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	GetAllUser(ctx context.Context) ([]User, error)
}

// Compile-time check to make sure *Queries implement interface Querier
var _ Querier = (*Queries)(nil)
