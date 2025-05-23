package middleware

import (
	"net/http"
	"strings"

	"github.com/spaghetti-lover/qairlines/pkg/token"
	"github.com/spaghetti-lover/qairlines/pkg/utils"
)

func AuthMiddleware(tokenMaker token.Maker) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				utils.WriteError(w, http.StatusUnauthorized, "Authorization header is required", nil)
				return
			}

			// Kiểm tra định dạng token
			fields := strings.Fields(authHeader)
			if len(fields) != 2 || strings.ToLower(fields[0]) != "bearer" {
				utils.WriteError(w, http.StatusUnauthorized, "Invalid authorization format", nil)
				return
			}

			accessToken := fields[1]

			// Xác thực token
			payload, err := tokenMaker.VerifyToken(accessToken, 1)
			if err != nil {
				http.Error(w, `{"message": "Authentication failed. Access token required."}`, http.StatusUnauthorized)
				return
			}

			r = r.WithContext(utils.ContextWithUserId(r.Context(), payload.UserId))

			// Tiếp tục xử lý request
			next.ServeHTTP(w, r)
		})
	}
}
