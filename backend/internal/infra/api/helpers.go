package api

import (
	"encoding/json"
	"net/http"
)

// errorResponse format error reponse for http request
func errorResponse(w http.ResponseWriter, err error, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	payload := map[string]string{"error": err.Error()}
	if encodeErr := json.NewEncoder(w).Encode(payload); encodeErr != nil {
		http.Error(w, `{"error":"Internal server error"}`, http.StatusInternalServerError)
	}
}
