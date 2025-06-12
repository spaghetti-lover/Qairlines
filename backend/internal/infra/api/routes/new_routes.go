package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/handlers"
)

func RegisterNewsRoutes(router *gin.RouterGroup, newsHandler *handlers.NewsHandler) {
	new := router.Group("/news")
	{
		new.GET("/all", newsHandler.GetAllNews)
		new.GET("/", newsHandler.GetNews)
		new.DELETE("/", newsHandler.DeleteNews)
		new.POST("/", newsHandler.CreateNews)
	}
}
