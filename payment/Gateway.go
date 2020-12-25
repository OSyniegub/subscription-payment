package payment

import (
	"github.com/OSyniegub/subscription-payment/payment/dto"
)

type Gateway interface {
	PaymentIntent(amount int64) (string, error)
	PaymentConfirm(requestDto dto.PaymentConfirmRequestDto) (dto.PaymentConfirmResponseDto, error)
	CardTokenGenerate(requestDto dto.CardTokenGenerateRequestDto) (string, error)
}

func MakePaymentIntent(gateway Gateway, amount int64) (string, error)  {
	return gateway.PaymentIntent(amount)
}

func MakePaymentConfirm(gateway Gateway, requestDto dto.PaymentConfirmRequestDto) (dto.PaymentConfirmResponseDto, error)  {
	return gateway.PaymentConfirm(requestDto)
}

func MakeCardTokenGenerate(gateway Gateway, requestDto dto.CardTokenGenerateRequestDto) (string, error) {
	return gateway.CardTokenGenerate(requestDto)
}