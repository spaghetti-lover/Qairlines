package api

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
	"github.com/spaghetti-lover/qairlines/config"
	db "github.com/spaghetti-lover/qairlines/db/sqlc"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/auth"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/news"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/user"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/handlers"
	"github.com/spaghetti-lover/qairlines/internal/infra/postgresql"
	"github.com/spaghetti-lover/qairlines/pkg/token"
)

type Server struct {
	store      *db.Store
	router     http.Handler
	tokenMaker token.Maker
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config config.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	//Health
	healthRepo := postgresql.NewHealthRepositoryPostgres(store)
	healthUseCase := usecases.NewHealthUseCase(healthRepo)
	healthHandler := handlers.NewHealthHandler(healthUseCase)

	// User
	userRepo := postgresql.NewUserRepositoryPostgres(store, tokenMaker)
	userGetAllUseCase := user.NewUserGetAllUseCase(userRepo)
	userCreateUseCase := user.NewUserCreateUseCase(userRepo)
	userGetByEmailUseCase := user.NewUserGetByEmailUseCase(userRepo)
	userHandler := handlers.NewUserHandler(userGetAllUseCase, userCreateUseCase, userGetByEmailUseCase)

	//Auth
	loginUseCase := auth.NewLoginUseCase(userRepo, tokenMaker)
	authHandler := handlers.NewAuthHandler(loginUseCase)

	newsRepo := postgresql.NewNewsModelRepositoryPostgres(store)
	newsGetAllUseCase := news.NewNewsGetAllUseCase(newsRepo)
	newsHandler := handlers.NewNewsHandler(newsGetAllUseCase)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", healthHandler.ServeHTTP)

	// News api group
	mux.HandleFunc("GET /api/news", newsHandler.GetAllNews)

	// User api group
	mux.HandleFunc("POST /api/user", userHandler.CreateUser)

	mux.HandleFunc("POST /api/auth/login", authHandler.Login)

	// Booking api group

	mux.HandleFunc("GET /api/booking/{booking_id}", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("PUT /api/booking/{booking_id}", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("DELETE /api/booking/{booking_id}", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("GET /api/booking/passengers/{booking_id}", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("GET /api/booking", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("POST /api/booking", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("GET /api/booking/flight/{flight_id}", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("POST /api/booking/cancel/{booking_id}", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("GET /api/booking/info/{booking_id}", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	// Airplane api group
	mux.HandleFunc("POST /api/airplanes/models", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("GET /api/airplanes/models/{airplane_model_id}", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("PUT /api/airplanes/models/{airplane_model_id}", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("DELETE /api/airplanes/models/{airplane_model_id}", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("POST /api/airplanes", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("GET /api/airplanes/{airplane_id}", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("PUT /api/airplanes/{airplane_id}", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("DELETE /api/airplanes/{airplane_id}", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("GET /api/airplanes/by-regis/{registration_number}", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("GET /api/airplanes", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	// Airport api group

	mux.HandleFunc("GET /api/airports", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("POST /api/airports", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("GET /api/airports/{airport_id}", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("PUT /api/airports/{airport_id}", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("DELETE /api/airports/{airport_id}", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	// Flight api group
	mux.HandleFunc("GET /api/flights/{airport_id}", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("PUT /api/flights/{flight_id}", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("DELETE /api/flights/{flight_id}", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("GET /api/flights", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("POST /api/flights", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("GET /api/flights/search", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("GET /api/flights/passengers/{flight_id}", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("GET /api/flights/passengers/citizen/{citizen_id}", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("POST /api/flights/delay", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("GET /api/flights/flight-seats/{flight_id}", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("GET /api/flights/flight-seats-available/{flight_id}", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("GET /api/flights/flight-seats/{flight_id}/prices", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	// Payment api group
	mux.HandleFunc("GET /api/advert/{advert_name}", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("PUT /api/advert/{advert_name}", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("DELETE /api/advert/{advert_name}", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("GET /api/advert", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("POST /api/advert", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Cho phép tất cả các origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(mux)

	server := &Server{
		store:      store,
		router:     corsHandler,
		tokenMaker: tokenMaker,
	}

	return server, nil
}

// ServeHTTP satisfies http.Handler interface, so Server can be passed to http.ListenAndServe directly
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

func methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
