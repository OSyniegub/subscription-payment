package payment

import "github.com/OSyniegub/subscription-payment/payment/dto"

type Gateway interface {
	PaymentIntent(amount int64) (string, error)
	PaymentConfirm(clientSecret string) (dto.PaymentConfirmResponseDto, error)
}

func MakePaymentIntent(gateway Gateway, amount int64) (string, error)  {
	return gateway.PaymentIntent(amount)
}

func MakePaymentConfirm(gateway Gateway, paymentId string) (dto.PaymentConfirmResponseDto, error)  {
	return gateway.PaymentConfirm(paymentId)
}