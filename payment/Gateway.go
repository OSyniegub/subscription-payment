package payment

import "github.com/OSyniegub/subscription-payment/payment/dto"

type Gateway interface {
	PaymentIntent(amount string) (string, error)
	PaymentConfirm(requestDto dto.PaymentConfirmRequestDto) ([]byte, error)
	CardTokenGenerate(requestDto dto.CardTokenGenerateRequestDto) (string, error)
}
