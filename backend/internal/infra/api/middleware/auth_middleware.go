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
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is required"})
			c.Abort()
			return
		}

		fields := strings.Fields(authHeader)
		if len(fields) != 2 || strings.ToLower(fields[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization format"})
			c.Abort()
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken, token.TokenTypeAccessToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Invalid token"})
			c.Abort()
			return
		}

		// Lưu thông tin xác thực vào context để dùng ở handler sau
		c.Set(AuthorizationPayloadKey, payload)
		c.Next()
	}
}
