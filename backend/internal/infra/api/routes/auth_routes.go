package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/handlers"
)

func RegisterAuthRoutes(router *gin.RouterGroup, authHandler *handlers.AuthHandler) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)

		auth.PUT("/:id/password", authHandler.ChangePassword)

		auth.PUT("/change-password", authHandler.ChangePassword)
	}
}
