package gateway

type PaymentGateway interface {
	PaymentIntent(amount int64) (string, error)
	PaymentConfirm(clientSecret string) (string, error)
	ChargeGet(chargeId string) (string, error)
}