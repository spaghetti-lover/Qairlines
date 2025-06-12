package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
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
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/user"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/handlers"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/routes"
	"github.com/spaghetti-lover/qairlines/internal/infra/postgresql"
	"github.com/spaghetti-lover/qairlines/pkg/token"
)

type Server struct {
	store      *db.Store
	router     http.Handler
	tokenMaker token.Maker
}

func NewServer(config config.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	// Health
	healthRepo := postgresql.NewHealthRepositoryPostgres(store)
	customerRepo := postgresql.NewCustomerRepositoryPostgres(store, tokenMaker)
	userRepo := postgresql.NewUserRepositoryPostgres(store, tokenMaker)
	newsRepo := postgresql.NewNewsModelRepositoryPostgres(store)
	adminRepo := postgresql.NewAdminRepositoryPostgres(store, tokenMaker)
	flightRepo := postgresql.NewFlightRepositoryPostgres(store)
	ticketRepo := postgresql.NewTicketRepositoryPostgres(store)
	bookingRepo := postgresql.NewBookingRepositoryPostgres(store)

	healthUseCase := usecases.NewHealthUseCase(healthRepo)
	userUpdateUseCase := user.NewUserUpdateUseCase(userRepo)
	customerCreateUseCase := customer.NewCreateCustomerUseCase(customerRepo, userRepo)
	customerUpdateUseCase := customer.NewCustomerUpdateUseCase(customerRepo)
	customerGetAllUseCase := customer.NewGetAllCustomersUseCase(customerRepo)
	customerDeleteUseCase := customer.NewDeleteCustomerUseCase(customerRepo)
	customerGetUseCase := customer.NewGetCustomerDetailsUseCase(customerRepo, tokenMaker)
	loginUseCase := auth.NewLoginUseCase(userRepo, tokenMaker)
	changePasswordUseCase := auth.NewChangePasswordUseCase(userRepo)
	newsGetAllWithAuthorUseCase := news.NewNewsGetAllWithAuthorUseCase(newsRepo)
	newsGetUseCase := news.NewGetNewsUseCase(newsRepo)
	newsDeleteUseCase := news.NewDeleteNewsUseCase(newsRepo)
	newsCreateUseCase := news.NewCreateNewsUseCase(newsRepo)
	newsUpdateUseCase := news.NewUpdateNewsUseCase(newsRepo)
	adminCreateUseCase := admin.NewCreateAdminUseCase(adminRepo, userRepo)
	getAllAdminsUseCase := admin.NewGetAllAdminsUseCase(adminRepo)
	updateAdminUseCase := admin.NewUpdateAdminUseCase(adminRepo, userRepo)
	getCurrentAdminUseCase := admin.NewGetCurrentAdminUseCase(adminRepo)
	deleteAdminUseCase := admin.NewDeleteAdminUseCase(adminRepo)
	flightCreateUseCase := flight.NewCreateFlightUseCase(flightRepo)
	flightGetUseCase := flight.NewGetFlightUseCase(flightRepo)
	flightUpdateUseCase := flight.NewUpdateFlightTimesUseCase(flightRepo)
	flightGetAllUseCase := flight.NewGetAllFlightsUseCase(flightRepo, ticketRepo)
	flightDeleteUseCase := flight.NewDeleteFlightUseCase(flightRepo)
	flightSearchUseCase := flight.NewSearchFlightsUseCase(flightRepo)
	flightSuggestedUseCase := flight.NewGetSuggestedFlightsUseCase(flightRepo)
	ticketGetTicketByFlightIDUseCase := ticket.NewGetTicketsByFlightIDUseCase(ticketRepo)
	ticketCancelUseCase := ticket.NewCancelTicketUseCase(ticketRepo)
	ticketGetUseCase := ticket.NewGetTicketUseCase(ticketRepo)
	ticketUpdateUseCase := ticket.NewUpdateSeatsUseCase(ticketRepo)
	bookingCreateUseCase := booking.NewCreateBookingUseCase(bookingRepo, flightRepo)
	bookingGetUseCase := booking.NewGetBookingUseCase(bookingRepo)

	healthHandler := handlers.NewHealthHandler(healthUseCase)
	customerHandler := handlers.NewCustomerHandler(customerCreateUseCase, customerUpdateUseCase, userUpdateUseCase, customerGetAllUseCase, customerDeleteUseCase, customerGetUseCase)
	authHandler := handlers.NewAuthHandler(loginUseCase, changePasswordUseCase)
	newsHandler := handlers.NewNewsHandler(newsGetAllWithAuthorUseCase, newsDeleteUseCase, newsCreateUseCase, newsUpdateUseCase, newsGetUseCase)
	adminHandler := handlers.NewAdminHandler(adminCreateUseCase, getCurrentAdminUseCase, getAllAdminsUseCase, updateAdminUseCase, deleteAdminUseCase)
	flightHandler := handlers.NewFlightHandler(flightCreateUseCase, flightGetUseCase, flightUpdateUseCase, flightGetAllUseCase, flightDeleteUseCase, flightSearchUseCase, flightSuggestedUseCase)
	ticketHandler := handlers.NewTicketHandler(ticketGetTicketByFlightIDUseCase, ticketGetUseCase, ticketCancelUseCase, ticketUpdateUseCase)
	bookingHandler := handlers.NewBookingHandler(bookingCreateUseCase, tokenMaker, userRepo, bookingGetUseCase)
	// Middleware
	//authMiddleware := middleware.AuthMiddleware(tokenMaker)

	// Create a new Gin router
	router := gin.Default()

	// Group all APIs under "/api"
	apiRouter := router.Group("/api")

	// Health API
	router.GET("/health", healthHandler.GetHealth)

	// News API
	routes.RegisterNewsRoutes(apiRouter, newsHandler)

	// Customer API
	routes.RegisterCustomerRoutes(apiRouter, customerHandler)

	// Auth API
	routes.RegisterAuthRoutes(apiRouter, authHandler)

	// Admin API
	routes.RegisterAdminRoutes(apiRouter, adminHandler)

	// Flight API
	routes.RegisterFlightRoutes(apiRouter, flightHandler)

	// Ticket API
	routes.RegisterTicketRoutes(apiRouter, ticketHandler)

	// Booking API
	routes.RegisterBookingRoutes(apiRouter, bookingHandler)

	// Statistic API
	routes.RegisterStatisticRoutes(apiRouter)

	// Wrap router with CORS middleware
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}).Handler(router)

	server := &Server{
		store:      store,
		router:     corsHandler,
		tokenMaker: tokenMaker,
	}

	return server, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
