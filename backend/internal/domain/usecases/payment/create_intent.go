package payment

import (
	"context"
	"fmt"

	"github.com/spaghetti-lover/qairlines/internal/domain/adapters"
)

type ICreatePaymentIntentUsecase interface {
	Execute(ctx context.Context, bookingID int64, amount int64, currency string) (string, error)
}

type CreatePaymentIntentUseCase struct {
	gateway adapters.PaymentGateway
}

func NewCreatePaymentIntentUseCase(gateway adapters.PaymentGateway) ICreatePaymentIntentUsecase {
	return &CreatePaymentIntentUseCase{gateway: gateway}
}

func (u *CreatePaymentIntentUseCase) Execute(ctx context.Context, bookingID int64, amount int64, currency string) (string, error) {
	metadata := map[string]string{"booking_id": fmt.Sprintf("%d", bookingID)}
	clientSecret, err := u.gateway.CreatePaymentIntent(amount, currency, metadata)
	if err != nil {
		return "", err
	}
	return clientSecret, nil
}
