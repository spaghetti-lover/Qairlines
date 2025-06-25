package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterStatisticRoutes(router *gin.RouterGroup) {
	router.GET("/statistic", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Statistics retrieved successfully.",
			"data": gin.H{
				"flights": 120,
				"tickets": 450,
				"revenue": 1145430000,
			},
		})
	})
}
