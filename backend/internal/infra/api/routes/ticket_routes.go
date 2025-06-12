package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/handlers"
)

func RegisterTicketRoutes(router *gin.RouterGroup, ticketHandler *handlers.TicketHandler) {
	ticket := router.Group("/ticket")
	{
		ticket.GET("/list", ticketHandler.GetTicketsByFlightID)
		ticket.PUT("/cancel", ticketHandler.CancelTicket)
		ticket.GET("/", ticketHandler.GetTicket)
		ticket.PUT("/update-seats", ticketHandler.UpdateSeats)
	}
}
