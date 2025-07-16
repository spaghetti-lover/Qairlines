package stripe

import (
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentintent"
)

type StripeGateway struct{}

func NewStripeGateway(secretKey string) *StripeGateway {
	stripe.Key = secretKey
	return &StripeGateway{}
}

func (s *StripeGateway) CreatePaymentIntent(amount int64, currency string, metadata map[string]string) (string, error) {
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String(currency),
	}
	for k, v := range metadata {
		params.AddMetadata(k, v)
	}
	intent, err := paymentintent.New(params)
	if err != nil {
		return "", err
	}
	return intent.ClientSecret, nil
}
