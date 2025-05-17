package http

import (
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
	server.router.HandleFunc("/health", server.handleHealth)
	return server, nil
}

// ServeHTTP satisfies http.Handler interface, so Server can be passed to http.ListenAndServe directly
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// Simple health check handler
func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
