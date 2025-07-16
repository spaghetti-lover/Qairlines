package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/handlers"
)

func RegisterPaymentRoutes(router *gin.RouterGroup, paymentHandler *handlers.PaymentHandler) {
	payment := router.Group("/")
	payment.POST("/payment-intents", paymentHandler.CreatePaymentIntent)
}
