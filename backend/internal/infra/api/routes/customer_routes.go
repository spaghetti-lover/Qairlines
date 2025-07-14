package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/handlers"
)

func RegisterCustomerRoutes(router *gin.RouterGroup, customerHandler *handlers.CustomerHandler) {
	customer := router.Group("/customer")
	{
		customer.POST("/", customerHandler.CreateCustomerTx)
		customer.PUT("/:id", customerHandler.UpdateCustomer)
		customer.GET("", customerHandler.ListCustomers)
		customer.DELETE("/delete", customerHandler.DeleteCustomer)
		customer.GET("/", customerHandler.GetCustomerDetails)
	}
}
