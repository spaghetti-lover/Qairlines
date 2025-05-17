package api

import (
	"encoding/json"
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

// 
func errorResponse(w http.ResponseWriter, err error, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	payload := map[string]string{"error": err.Error()}
	if encodeErr := json.NewEncoder(w).Encode(payload); encodeErr != nil {
		http.Error(w, `{"error":"Internal server error"}`, http.StatusInternalServerError)
	}
}
