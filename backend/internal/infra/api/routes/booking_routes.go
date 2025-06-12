package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/handlers"
)

func RegisterBookingRoutes(router *gin.RouterGroup, bookingHandler *handlers.BookingHandler) {
	booking := router.Group("/booking")
	{
		booking.POST("/", bookingHandler.CreateBooking)
		booking.GET("/", bookingHandler.GetBooking)
	}
}
