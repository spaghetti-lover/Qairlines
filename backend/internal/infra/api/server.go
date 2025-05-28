package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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
	"github.com/spaghetti-lover/qairlines/internal/infra/api/middleware"
	"github.com/spaghetti-lover/qairlines/internal/infra/postgresql"
	"github.com/spaghetti-lover/qairlines/pkg/token"
	"github.com/spaghetti-lover/qairlines/pkg/utils"
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
	flightGetAllUseCase := flight.NewGetAllFlightsUseCase(flightRepo)
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
	authMiddleware := middleware.AuthMiddleware(tokenMaker)

	// Use gorilla/mux for routing
	router := mux.NewRouter()

	// Group all APIs under "/api"
	apiRouter := router.PathPrefix("/api").Subrouter()

	// Health API
	router.HandleFunc("/health", healthHandler.ServeHTTP).Methods("GET")

	// News API
	apiRouter.HandleFunc("/news/all", newsHandler.GetAllNews).Methods("GET")
	apiRouter.Handle("/news", authMiddleware(http.HandlerFunc(newsHandler.GetNews))).Methods("GET")
	apiRouter.Handle("/news", authMiddleware(http.HandlerFunc(newsHandler.DeleteNews))).Methods("DELETE")
	apiRouter.Handle("/news", authMiddleware(http.HandlerFunc(newsHandler.CreateNews))).Methods("POST")
	apiRouter.Handle("/news", authMiddleware(http.HandlerFunc(newsHandler.UpdateNews))).Methods("PUT")

	// Customer API
	apiRouter.HandleFunc("/customer", customerHandler.CreateCustomerTx).Methods("POST")
	apiRouter.Handle("/customer/{id}", authMiddleware(http.HandlerFunc(customerHandler.UpdateCustomer))).Methods("PUT")
	apiRouter.Handle("/customer/all", authMiddleware(http.HandlerFunc(customerHandler.GetAllCustomers))).Methods("GET")
	apiRouter.Handle("/customer/delete", authMiddleware(http.HandlerFunc(customerHandler.DeleteCustomer))).Methods("DELETE")
	apiRouter.Handle("/customer", authMiddleware(http.HandlerFunc(customerHandler.GetCustomerDetails))).Methods("GET")

	// Auth API
	apiRouter.HandleFunc("/auth/login", authHandler.Login).Methods("POST")
	apiRouter.Handle("/change-password", authMiddleware(http.HandlerFunc(authHandler.ChangePassword))).Methods("PUT")
	apiRouter.Handle("/auth/{id}/password", authMiddleware(http.HandlerFunc(authHandler.ChangePassword))).Methods("PUT")

	// Admin API
	apiRouter.Handle("/admin", authMiddleware(http.HandlerFunc(adminHandler.GetCurrentAdmin))).Methods("GET")
	apiRouter.HandleFunc("/admin", adminHandler.CreateAdminTx).Methods("POST")
	apiRouter.Handle("/admin/all", authMiddleware(http.HandlerFunc(adminHandler.GetAllAdmins))).Methods("GET")
	apiRouter.Handle("/admin", authMiddleware(http.HandlerFunc(adminHandler.UpdateAdmin))).Methods("PUT")
	apiRouter.Handle("/admin", authMiddleware(http.HandlerFunc(adminHandler.DeleteAdmin))).Methods("DELETE")

	// Flight API
	apiRouter.Handle("/flight", authMiddleware(http.HandlerFunc(flightHandler.CreateFlight))).Methods("POST")
	apiRouter.HandleFunc("/flight", flightHandler.GetFlight).Methods("GET")
	apiRouter.Handle("/flight/update", authMiddleware(http.HandlerFunc(flightHandler.UpdateFlightTimes))).Methods("PUT")
	apiRouter.Handle("/flight/all", authMiddleware(http.HandlerFunc(flightHandler.GetAllFlights))).Methods("GET")
	apiRouter.Handle("/flight", authMiddleware(http.HandlerFunc(flightHandler.DeleteFlight))).Methods("DELETE")
	apiRouter.HandleFunc("/flight/search", flightHandler.SearchFlights).Methods("GET")
	apiRouter.HandleFunc("/flight/suggest", flightHandler.GetSuggestedFlights).Methods("GET")

	// Ticket API
	apiRouter.Handle("/ticket/list", authMiddleware(http.HandlerFunc(ticketHandler.GetTicketsByFlightID))).Methods("GET")
	apiRouter.Handle("/ticket/cancel", authMiddleware(http.HandlerFunc(ticketHandler.CancelTicket))).Methods("PUT")
	apiRouter.Handle("/ticket", authMiddleware(http.HandlerFunc(ticketHandler.GetTicket))).Methods("GET")
	apiRouter.Handle("/ticket/update-seats", authMiddleware(http.HandlerFunc(ticketHandler.UpdateSeats))).Methods("PUT")

	// Booking API
	apiRouter.Handle("/booking", authMiddleware(http.HandlerFunc(bookingHandler.CreateBooking))).Methods("POST")
	// Statistic API
	apiRouter.HandleFunc("/statistic", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"message": "Statistics retrieved successfully.",
			"data": map[string]interface{}{
				"flights": 120,
				"tickets": 450,
				"revenue": 1145430000,
			},
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			utils.WriteError(w, http.StatusInternalServerError, "failed to encode response", err)
		}
	}).Methods("GET")
	apiRouter.Handle("/booking", authMiddleware(http.HandlerFunc(bookingHandler.GetBooking))).Methods("GET")

	// Wrap router with CORS middleware
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
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
