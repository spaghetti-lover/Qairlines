package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
	"github.com/spaghetti-lover/qairlines/config"
	db "github.com/spaghetti-lover/qairlines/db/sqlc"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/di"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/middleware"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/routes"
	"github.com/spaghetti-lover/qairlines/pkg/token"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Server struct {
	store      *db.Store
	router     http.Handler
	tokenMaker token.Maker
}

func NewServer(config config.Config, store *db.Store) (*Server, error) {
	container, err := di.NewContainer(config, store)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize dependencies: %w", err)
	}
	httpLogger := zerolog.New(&lumberjack.Logger{
		Filename:   "logs/http.log",
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     5,    // days
		Compress:   true, // disabled by default
		LocalTime:  true,
	}).With().Timestamp().Logger()

	recoveryLogger := zerolog.New(&lumberjack.Logger{
		Filename:   "logs/recovery.log",
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     5,    // days
		Compress:   true, // disabled by default
		LocalTime:  true,
	}).With().Timestamp().Logger()

	// Create a new Gin router
	router := gin.Default()
	router.Use(middleware.RateLimitingMiddleware(), middleware.LoggerMiddleware(httpLogger), middleware.RecoveryMiddleware(recoveryLogger))

	gin.SetMode(gin.DebugMode)

	// Clean up clients for rate limiting
	go middleware.CleanUpClients()

	// Group all APIs under "/api"
	apiRouter := router.Group("/api")

	// Health API
	router.GET("/health", container.HealthHandler.GetHealth)
	// News API
	routes.RegisterNewsRoutes(apiRouter, container.NewsHandler)
	// Customer API
	routes.RegisterCustomerRoutes(apiRouter, container.CustomerHandler)
	// Auth API
	routes.RegisterAuthRoutes(apiRouter, container.AuthHandler)
	// Admin API
	routes.RegisterAdminRoutes(apiRouter, container.AdminHandler)
	// Flight API
	routes.RegisterFlightRoutes(apiRouter, container.FlightHandler)
	// Ticket API
	routes.RegisterTicketRoutes(apiRouter, container.TicketHandler)
	// Booking API
	routes.RegisterBookingRoutes(apiRouter, container.BookingHandler)
	// Statistic API
	routes.RegisterStatisticRoutes(apiRouter)
	// View Static File
	router.StaticFS("/images", gin.Dir("./uploads", false))

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
		tokenMaker: container.TokenMaker,
	}

	return server, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
