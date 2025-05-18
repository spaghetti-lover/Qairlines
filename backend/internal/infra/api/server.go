package api

import (
	"errors"
	"net/http"

	db "github.com/spaghetti-lover/qairlines/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *http.ServeMux
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(store *db.Store) (*Server, error) {
	server := &Server{
		store:  store,
		router: http.NewServeMux(),
	}
	server.router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})

	// User api group
	server.router.HandleFunc("GET /api/user/{user_id}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("PUT /api/user/{user_id}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("GET /api/user/", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("POST /api/user/", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("GET /api/user/username/{user_name}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("DELETE /api/user/{user_name}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("GET /api/user/me/", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("POST /api/user/auth", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("POST /api/user/mail", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})

	// Booking api group
	server.router.HandleFunc("GET /api/booking/{booking_id}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("PUT /api/booking/{booking_id}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("DELETE /api/booking/{booking_id}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("GET /api/booking/passengers/{booking_id}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("GET /api/booking", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("POST /api/booking", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("GET /api/booking/flight/{flight_id}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("POST /api/booking/cancel/{booking_id}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("GET /api/booking/info/{booking_id}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})

	// Airplane api group
	server.router.HandleFunc("POST /api/airplanes/models", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("GET /api/airplanes/models/{airplane_model_id}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("PUT /api/airplanes/models/{airplane_model_id}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("DELETE /api/airplanes/models/{airplane_model_id}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("POST /api/airplanes", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("GET /api/airplanes/{airplane_id}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("PUT /api/airplanes/{airplane_id}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("DELETE /api/airplanes/{airplane_id}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("GET /api/airplanes/by-regis/{registration_number}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("GET /api/airplanes", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})

	// Airport api group
	server.router.HandleFunc("GET /api/airports", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("POST /api/airports", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("GET /api/airports/{airport_id}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("PUT /api/airports/{airport_id}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("DELETE /api/airports/{airport_id}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})

	// Flight api group
	server.router.HandleFunc("GET /api/flights/{airport_id}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("PUT /api/flights/{flight_id}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("DELETE /api/flights/{flight_id}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("GET /api/flights", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("POST /api/flights", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("GET /api/flights/search", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("GET /api/flights/passengers/{flight_id}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("GET /api/flights/passengers/citizen/{citizen_id}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("POST /api/flights/delay", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("GET /api/flights/flight-seats/{flight_id}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("GET /api/flights/flight-seats-available/{flight_id}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("GET /api/flights/flight-seats/{flight_id}/prices", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})

	// Payment api group
	server.router.HandleFunc("GET /api/advert/{advert_name}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("PUT /api/advert/{advert_name}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("DELETE /api/advert/{advert_name}", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("GET /api/advert", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	server.router.HandleFunc("POST /api/advert", func(w http.ResponseWriter, r *http.Request) {
		errorResponse(w, errors.New("not implemented"), http.StatusInternalServerError)
	})
	return server, nil
}

// ServeHTTP satisfies http.Handler interface, so Server can be passed to http.ListenAndServe directly
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
