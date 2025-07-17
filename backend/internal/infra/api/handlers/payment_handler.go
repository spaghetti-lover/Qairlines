package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/payment"
)

type PaymentHandler struct {
	createPaymentUseCase payment.ICreatePaymentIntentUsecase
}

func NewPaymentHandler(createPaymentUseCase payment.ICreatePaymentIntentUsecase) *PaymentHandler {
	return &PaymentHandler{createPaymentUseCase: createPaymentUseCase}
}

func (h *PaymentHandler) CreatePaymentIntent(ctx *gin.Context) {
	var req struct {
		BookingID int64  `json:"booking_id"`
		Amount    int64  `json:"amount"`
		Currency  string `json:"currency"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	clientSecret, err := h.createPaymentUseCase.Execute(ctx, req.BookingID, req.Amount, req.Currency)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"client_secret": clientSecret})
}
