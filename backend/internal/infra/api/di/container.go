package di

import (
	"github.com/spaghetti-lover/qairlines/config"
	db "github.com/spaghetti-lover/qairlines/db/sqlc"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/admin"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/auth"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/booking"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/customer"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/flight"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/news"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/ticket"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/handlers"
	"github.com/spaghetti-lover/qairlines/internal/infra/postgresql"
	"github.com/spaghetti-lover/qairlines/pkg/token"
)

type Container struct {
	HealthHandler   *handlers.HealthHandler
	CustomerHandler *handlers.CustomerHandler
	AuthHandler     *handlers.AuthHandler
	NewsHandler     *handlers.NewsHandler
	AdminHandler    *handlers.AdminHandler
	FlightHandler   *handlers.FlightHandler
	TicketHandler   *handlers.TicketHandler
	BookingHandler  *handlers.BookingHandler
	TokenMaker      token.Maker
}

func NewContainer(config config.Config, store *db.Store) (*Container, error) {
	// Token Maker
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}

	// Repositories
	healthRepo := postgresql.NewHealthRepositoryPostgres(store)
	customerRepo := postgresql.NewCustomerRepositoryPostgres(store, tokenMaker)
	userRepo := postgresql.NewUserRepositoryPostgres(store, tokenMaker)
	newsRepo := postgresql.NewNewsModelRepositoryPostgres(store)
	adminRepo := postgresql.NewAdminRepositoryPostgres(store, tokenMaker)
	flightRepo := postgresql.NewFlightRepositoryPostgres(store)
	ticketRepo := postgresql.NewTicketRepositoryPostgres(store)
	bookingRepo := postgresql.NewBookingRepositoryPostgres(store)

	// Use Cases
	healthUseCase := usecases.NewHealthUseCase(healthRepo)
	customerCreateUseCase := customer.NewCreateCustomerUseCase(customerRepo, userRepo)
	customerUpdateUseCase := customer.NewCustomerUpdateUseCase(customerRepo)
	customerListAllUseCase := customer.NewListCustomersUseCase(customerRepo)
	customerDeleteUseCase := customer.NewDeleteCustomerUseCase(customerRepo)
	customerGetUseCase := customer.NewGetCustomerDetailsUseCase(customerRepo, tokenMaker)
	loginUseCase := auth.NewLoginUseCase(userRepo, tokenMaker)
	changePasswordUseCase := auth.NewChangePasswordUseCase(userRepo)
	newsGetAllWithAuthorUseCase := news.NewListNewsUseCase(newsRepo)
	newsGetUseCase := news.NewGetNewsUseCase(newsRepo)
	newsDeleteUseCase := news.NewDeleteNewsUseCase(newsRepo)
	newsCreateUseCase := news.NewCreateNewsUseCase(newsRepo)
	newsUpdateUseCase := news.NewUpdateNewsUseCase(newsRepo)
	adminCreateUseCase := admin.NewCreateAdminUseCase(adminRepo, userRepo)
	ListAdminsUseCase := admin.NewListAdminsUseCase(adminRepo)
	updateAdminUseCase := admin.NewUpdateAdminUseCase(adminRepo, userRepo)
	getCurrentAdminUseCase := admin.NewGetCurrentAdminUseCase(adminRepo)
	deleteAdminUseCase := admin.NewDeleteAdminUseCase(adminRepo)
	flightCreateUseCase := flight.NewCreateFlightUseCase(flightRepo)
	flightGetUseCase := flight.NewGetFlightUseCase(flightRepo)
	flightUpdateUseCase := flight.NewUpdateFlightTimesUseCase(flightRepo)
	flightGetAllUseCase := flight.NewGetAllFlightsUseCase(flightRepo, ticketRepo)
	flightDeleteUseCase := flight.NewDeleteFlightUseCase(flightRepo)
	flightSearchUseCase := flight.NewSearchFlightsUseCase(flightRepo)
	flightSuggestedUseCase := flight.NewlistFlightsUseCase(flightRepo)
	ticketGetTicketByFlightIDUseCase := ticket.NewGetTicketsByFlightIDUseCase(ticketRepo)
	ticketCancelUseCase := ticket.NewCancelTicketUseCase(ticketRepo)
	ticketGetUseCase := ticket.NewGetTicketUseCase(ticketRepo)
	ticketUpdateUseCase := ticket.NewUpdateSeatsUseCase(ticketRepo)
	bookingCreateUseCase := booking.NewCreateBookingUseCase(bookingRepo, flightRepo)
	bookingGetUseCase := booking.NewGetBookingUseCase(bookingRepo)

	// Handlers
	healthHandler := handlers.NewHealthHandler(healthUseCase)
	customerHandler := handlers.NewCustomerHandler(customerCreateUseCase, customerUpdateUseCase, nil, customerListAllUseCase, customerDeleteUseCase, customerGetUseCase)
	authHandler := handlers.NewAuthHandler(loginUseCase, changePasswordUseCase)
	newsHandler := handlers.NewNewsHandler(newsGetAllWithAuthorUseCase, newsDeleteUseCase, newsCreateUseCase, newsUpdateUseCase, newsGetUseCase, &config)
	adminHandler := handlers.NewAdminHandler(adminCreateUseCase, getCurrentAdminUseCase, ListAdminsUseCase, updateAdminUseCase, deleteAdminUseCase)
	flightHandler := handlers.NewFlightHandler(flightCreateUseCase, flightGetUseCase, flightUpdateUseCase, flightGetAllUseCase, flightDeleteUseCase, flightSearchUseCase, flightSuggestedUseCase)
	ticketHandler := handlers.NewTicketHandler(ticketGetTicketByFlightIDUseCase, ticketGetUseCase, ticketCancelUseCase, ticketUpdateUseCase)
	bookingHandler := handlers.NewBookingHandler(bookingCreateUseCase, tokenMaker, userRepo, bookingGetUseCase)

	return &Container{
		HealthHandler:   healthHandler,
		CustomerHandler: customerHandler,
		AuthHandler:     authHandler,
		NewsHandler:     newsHandler,
		AdminHandler:    adminHandler,
		FlightHandler:   flightHandler,
		TicketHandler:   ticketHandler,
		BookingHandler:  bookingHandler,
		TokenMaker:      tokenMaker,
	}, nil
}
