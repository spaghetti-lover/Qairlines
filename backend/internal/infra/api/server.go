package api

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/rs/cors"
	"github.com/spaghetti-lover/qairlines/config"
	db "github.com/spaghetti-lover/qairlines/db/sqlc"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/di"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/middleware"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/routes"
	"github.com/spaghetti-lover/qairlines/internal/infra/worker"
	"github.com/spaghetti-lover/qairlines/pkg/logger"
	"github.com/spaghetti-lover/qairlines/pkg/token"
)

type Server struct {
	store           db.Store
	router          http.Handler
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

func NewServer(config config.Config, store db.Store, redis *redis.Client, taskDistributor worker.TaskDistributor) (*Server, error) {
	container, err := di.NewContainer(config, redis, &store, taskDistributor)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize dependencies: %w", err)
	}

	httpLogger := logger.NewLoggerWithPath("logs/http.log", "info")
	recoveryLogger := logger.NewLoggerWithPath("logs/recovery.log", "warning")
	rateLimiterLogger := logger.NewLoggerWithPath("logs/rate_limiter.log", "warning")

	// Create a new Gin router
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(middleware.RateLimitingMiddleware(rateLimiterLogger), middleware.TraceMiddleware(), middleware.LoggerMiddleware(httpLogger), middleware.RecoveryMiddleware(recoveryLogger))

	gin.SetMode(gin.TestMode)

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
	// Payment API
	routes.RegisterPaymentRoutes(apiRouter, container.PaymentHandler)

	// Wrap router with CORS middleware
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}).Handler(router)

	server := &Server{
		store:           store,
		router:          corsHandler,
		tokenMaker:      container.TokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
