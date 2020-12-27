package payment

import "github.com/OSyniegub/subscription-payment/payment/dto"

type GooglePay struct {}

//TODO implement separate GooglePay payment processing
//For now it uses Stripe (card token generates on the frontend side from google api)
//and then replaced with stripe card token

func (s GooglePay) PaymentIntent(amount string) (string, error) {
	return "", nil
}

func (s GooglePay) PaymentConfirm(requestDto dto.PaymentConfirmRequestDto) ([]byte, error) {
	return []byte(""), nil
}

func (s GooglePay) CardTokenGenerate(requestDto dto.CardTokenGenerateRequestDto) (string, error) {
	return "", nil
}
