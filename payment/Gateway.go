package payment

import "github.com/OSyniegub/subscription-payment/payment/dto"

type Gateway interface {
	PaymentIntent(amount string) (string, error)
	PaymentConfirm(requestDto dto.PaymentConfirmRequestDto) ([]byte, error)
	CardTokenGenerate(requestDto dto.CardTokenGenerateRequestDto) (string, error)
}

func MakePaymentIntent(gateway Gateway, amount string) (string, error)  {
	return gateway.PaymentIntent(amount)
}

func MakePaymentConfirm(gateway Gateway, requestDto dto.PaymentConfirmRequestDto) ([]byte, error)  {
	return gateway.PaymentConfirm(requestDto)
}

func MakeCardTokenGenerate(gateway Gateway, requestDto dto.CardTokenGenerateRequestDto) (string, error) {
	return gateway.CardTokenGenerate(requestDto)
}