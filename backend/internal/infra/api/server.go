package api

import (
	"net/http"
	"strings"

	db "github.com/spaghetti-lover/qairlines/db/sqlc"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/handlers"
	"github.com/spaghetti-lover/qairlines/internal/infra/postgresql"
)

type Server struct {
	store  *db.Store
	router *http.ServeMux
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(store *db.Store) (*Server, error) {
	healthRepo := postgresql.NewHealthRepositoryPostgres(store)
	healthUseCase := usecases.NewHealthUseCase(healthRepo)
	healthHandler := handlers.NewHealthHandler(healthUseCase)



	server := &Server{
		store:  store,
		router: http.NewServeMux(),
	}
	server.router.Handle("/health", withMethod("GET", healthHandler.ServeHTTP))

	server.router.Handle("/api", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/api")
		switch {
		// --- USER ---
		case path == "/user" && r.Method == http.MethodGet:
			notImplemented(w, r)
		case path == "/user" && r.Method == http.MethodPost:
			notImplemented(w, r)
		case path == "/user/me" && r.Method == http.MethodGet:
			notImplemented(w, r)
		case strings.HasPrefix(path, "/user/username/") && r.Method == http.MethodGet:
			notImplemented(w, r)
		case strings.Count(path, "/") == 2 && strings.HasPrefix(path, "/user/") && r.Method == http.MethodGet:
			// GET /api/user/{user_id}
			notImplemented(w, r)
		case strings.Count(path, "/") == 2 && strings.HasPrefix(path, "/user/") && r.Method == http.MethodPut:
			notImplemented(w, r)
		case strings.Count(path, "/") == 2 && strings.HasPrefix(path, "/user/") && r.Method == http.MethodDelete:
			notImplemented(w, r)

		// --- BOOKING ---
		case path == "/booking" && r.Method == http.MethodGet:
			notImplemented(w, r)
		case path == "/booking" && r.Method == http.MethodPost:
			notImplemented(w, r)
		case strings.HasPrefix(path, "/booking/flight/") && r.Method == http.MethodGet:
			notImplemented(w, r)
		case strings.HasPrefix(path, "/booking/passengers/") && r.Method == http.MethodGet:
			notImplemented(w, r)
		case strings.HasPrefix(path, "/booking/info/") && r.Method == http.MethodGet:
			notImplemented(w, r)
		case strings.HasPrefix(path, "/booking/cancel/") && r.Method == http.MethodPost:
			notImplemented(w, r)
		case strings.Count(path, "/") == 2 && strings.HasPrefix(path, "/booking/") && r.Method == http.MethodGet:
			// GET /api/booking/{booking_id}
			notImplemented(w, r)
		case strings.Count(path, "/") == 2 && strings.HasPrefix(path, "/booking/") && r.Method == http.MethodPut:
			notImplemented(w, r)
		case strings.Count(path, "/") == 2 && strings.HasPrefix(path, "/booking/") && r.Method == http.MethodDelete:
			notImplemented(w, r)

		// --- AIRPLANES & MODELS ---
		case path == "/airplanes" && r.Method == http.MethodGet:
			notImplemented(w, r)
		case path == "/airplanes" && r.Method == http.MethodPost:
			notImplemented(w, r)
		case strings.HasPrefix(path, "/airplanes/by-regis/") && r.Method == http.MethodGet:
			notImplemented(w, r)

		case path == "/airplanes/models" && r.Method == http.MethodPost:
			notImplemented(w, r)
		case strings.Count(path, "/") == 3 && strings.HasPrefix(path, "/airplanes/models/"):
			// method branch
			switch r.Method {
			case http.MethodGet, http.MethodPut, http.MethodDelete:
				notImplemented(w, r)
			default:
				methodNotAllowed(w, r)
			}

		// --- AIRPORTS ---
		case path == "/airports" && (r.Method == http.MethodGet || r.Method == http.MethodPost):
			notImplemented(w, r)
		case strings.Count(path, "/") == 2 && strings.HasPrefix(path, "/airports/"):
			// GET, PUT, DELETE /api/airports/{airport_id}
			switch r.Method {
			case http.MethodGet, http.MethodPut, http.MethodDelete:
				notImplemented(w, r)
			default:
				methodNotAllowed(w, r)
			}

		// --- FLIGHTS ---
		case path == "/flights" && (r.Method == http.MethodGet || r.Method == http.MethodPost):
			notImplemented(w, r)
		case strings.HasPrefix(path, "/flights/search") && r.Method == http.MethodGet:
			notImplemented(w, r)
		case strings.HasPrefix(path, "/flights/passengers/citizen/") && r.Method == http.MethodGet:
			notImplemented(w, r)
		case strings.HasPrefix(path, "/flights/passengers/") && r.Method == http.MethodGet:
			notImplemented(w, r)
		case strings.HasPrefix(path, "/flights/delay") && r.Method == http.MethodPost:
			notImplemented(w, r)
		case strings.HasSuffix(path, "/flight-seats") && r.Method == http.MethodGet:
			notImplemented(w, r)
		case strings.HasSuffix(path, "/flight-seats-available") && r.Method == http.MethodGet:
			notImplemented(w, r)
		case strings.Contains(path, "/flight-seats/") && strings.HasSuffix(path, "/prices") && r.Method == http.MethodGet:
			notImplemented(w, r)
		case strings.Count(path, "/") == 2 && strings.HasPrefix(path, "/flights/") && (r.Method == http.MethodGet || r.Method == http.MethodPut || r.Method == http.MethodDelete):
			// /api/flights/{flight_id}
			notImplemented(w, r)

		// --- ADVERTS ---
		case path == "/advert" && (r.Method == http.MethodGet || r.Method == http.MethodPost):
			notImplemented(w, r)
		case strings.Count(path, "/") == 2 && strings.HasPrefix(path, "/advert/"):
			// GET, PUT, DELETE /api/advert/{advert_name}
			switch r.Method {
			case http.MethodGet, http.MethodPut, http.MethodDelete:
				notImplemented(w, r)
			default:
				methodNotAllowed(w, r)
			}

		default:
			http.NotFound(w, r)
		}
	}))
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
