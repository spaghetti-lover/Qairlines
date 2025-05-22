package api

import (
	"net/http"

	"github.com/rs/cors"
	db "github.com/spaghetti-lover/qairlines/db/sqlc"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/handlers"
	"github.com/spaghetti-lover/qairlines/internal/infra/postgresql"
)

type Server struct {
	store  *db.Store
	router http.Handler
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(store *db.Store) (*Server, error) {
	//Check Health
	healthRepo := postgresql.NewHealthRepositoryPostgres(store)
	healthUseCase := usecases.NewHealthUseCase(healthRepo)
	healthHandler := handlers.NewHealthHandler(healthUseCase)

	userRepo := postgresql.NewUserRepositoryPostgres(store)
	userGetAllUseCase := usecases.NewUserGetAllUseCase(userRepo)
	userGetAllHandler := handlers.NewUserGetHandler(userGetAllUseCase)

	newsRepo := postgresql.NewNewsModelRepositoryPostgres(store)
	newsGetAllUseCase := usecases.NewNewsGetAllUseCase(newsRepo)
	newsHandler := handlers.NewNewsHandler(newsGetAllUseCase)

	mux := http.NewServeMux()

	mux.Handle("/health", withMethod("GET", healthHandler.ServeHTTP))

	// News api group
	mux.Handle("/api/news", withMethod("GET", newsHandler.GetAllNews))

	// User api group
	mux.Handle("/api/user", withMethod("GET", userGetAllHandler.ServeHTTP))

	mux.HandleFunc("GET /api/user/{user_id}", func(w http.ResponseWriter, r *http.Request) {
		withMethod("GET", healthHandler.ServeHTTP)
	})

	mux.HandleFunc("PUT /api/user/{user_id}", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("POST /api/user/", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("GET /api/user/username/{user_name}", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("DELETE /api/user/{user_name}", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("GET /api/user/me/", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("POST /api/user/auth", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

	mux.HandleFunc("POST /api/user/mail", func(w http.ResponseWriter, r *http.Request) {
		notImplemented(w, r)
	})

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
		store:  store,
		router: corsHandler, // Gán corsHandler trực tiếp mà không cần ép kiểu
	}

	return server, nil
}

// ServeHTTP satisfies http.Handler interface, so Server can be passed to http.ListenAndServe directly
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func withMethod(method string, h http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			methodNotAllowed(w, r)
			return
		}
		h(w, r)
	})
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

func methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
