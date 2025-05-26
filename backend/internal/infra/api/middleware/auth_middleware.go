package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/spaghetti-lover/qairlines/pkg/token"
)

// Define a custom key type for context
type contextKey string

// Define the key used to store the authorization payload in the request context
const AuthorizationPayloadKey contextKey = "authorization_payload"

// AuthMiddleware creates a middleware for authorization
func AuthMiddleware(tokenMaker token.Maker) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, `{"message": "Authorization header is required"}`, http.StatusUnauthorized)
				return
			}

			// Check token format
			fields := strings.Fields(authHeader)
			if len(fields) != 2 || strings.ToLower(fields[0]) != "bearer" {
				http.Error(w, `{"message": "Invalid authorization format"}`, http.StatusUnauthorized)
				return
			}

			// Get access token
			accessToken := fields[1]

			// Verify token
			payload, err := tokenMaker.VerifyToken(accessToken, token.TokenTypeAccessToken)
			if err != nil {
				http.Error(w, `{"message": "Authentication failed. Invalid token."}`, http.StatusUnauthorized)
				return
			}

			// Add payload to context using the proper key
			ctx := context.WithValue(r.Context(), AuthorizationPayloadKey, payload)

			// Continue with the modified request
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
