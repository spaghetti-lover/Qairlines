package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/handlers"
)

func RegisterFlightRoutes(router *gin.RouterGroup, flightHandler *handlers.FlightHandler) {
	flight := router.Group("/flight")
	{
		flight.POST("/", flightHandler.CreateFlight)
		flight.GET("/", flightHandler.GetFlight)
		flight.PUT("/update", flightHandler.UpdateFlightTimes)
		flight.GET("/all", flightHandler.GetAllFlights)
		flight.DELETE("/", flightHandler.DeleteFlight)
		flight.GET("/search", flightHandler.SearchFlights)
		flight.GET("/suggest", flightHandler.GetSuggestedFlights)
	}
}
