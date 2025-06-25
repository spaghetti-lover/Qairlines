package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spaghetti-lover/qairlines/pkg/token"
)

// Define the key used to store the authorization payload in the request context
const AuthorizationPayloadKey = "authorization_payload"

// AuthMiddleware creates a middleware for authorization
func AuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is required"})
			ctx.Abort()
			return
		}

		fields := strings.Fields(authHeader)
		if len(fields) != 2 || strings.ToLower(fields[0]) != "bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization format"})
			ctx.Abort()
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken, token.TokenTypeAccessToken)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Invalid token"})
			ctx.Abort()
			return
		}

		// Lưu thông tin xác thực vào context để dùng ở handler sau
		ctx.Set(AuthorizationPayloadKey, payload)
		ctx.Next()
	}
}
