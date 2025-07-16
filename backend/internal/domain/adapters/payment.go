package adapters

type PaymentGateway interface {
	CreatePaymentIntent(amount int64, currency string, metadata map[string]string) (clientSecret string, err error)
}
