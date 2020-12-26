package payment

import (
	"errors"
	"github.com/OSyniegub/subscription-payment/payment/dto"
)

type GooglePay struct {}

func (s GooglePay) PaymentIntent(amount string) (string, error) {
	//TODO implement separate GooglePay payment processing
	//For now it uses Stripe (hardcoded)
	return "", errors.New("")
}

func (s GooglePay) PaymentConfirm(requestDto dto.PaymentConfirmRequestDto) ([]byte, error) {
	//TODO implement separate GooglePay payment processing
	//For now it uses Stripe (hardcoded)
	return []byte(""), errors.New("")
}

func (s GooglePay) CardTokenGenerate(requestDto dto.CardTokenGenerateRequestDto) (string, error) {
	//TODO implement separate GooglePay payment processing
	//For now it uses Stripe (hardcoded)
	return "", errors.New("")
}
