package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/handlers"
)

func RegisterAdminRoutes(router *gin.RouterGroup, adminHandler *handlers.AdminHandler) {
	admin := router.Group("/admin")
	{
		admin.GET("/", adminHandler.GetCurrentAdmin)
		admin.POST("/", adminHandler.CreateAdminTx)
		admin.GET("/all", adminHandler.GetAllAdmins)
		admin.PUT("/", adminHandler.UpdateAdmin)
		admin.DELETE("/", adminHandler.DeleteAdmin)
	}
}
