package payment

import (
	"encoding/json"
	"github.com/OSyniegub/subscription-payment/payment/dto"
	"github.com/stripe/stripe-go/v71"
	"strconv"
)

type StripePaymentIntent interface {
	New(params *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error)
	Confirm(id string, params *stripe.PaymentIntentConfirmParams) (*stripe.PaymentIntent, error)
}

type StripeToken interface {
	New(params *stripe.TokenParams) (*stripe.Token, error)
}

type Stripe struct {
	DoPaymentIntent StripePaymentIntent
	Token StripeToken
}

func (s Stripe) PaymentIntent(amount string) (string, error) {
	amountInt, err := strconv.Atoi(amount)

	if err != nil {
		return "", err
	}

	params := &stripe.PaymentIntentParams{
		Amount: stripe.Int64(int64(amountInt)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
	}

	params.AddMetadata("integration_check", "accept_a_payment")

	pi, err := s.DoPaymentIntent.New(params)

	if err != nil {
		return "", err
	}

	return pi.ID, err
}

func (s Stripe) PaymentConfirm(requestDto dto.PaymentConfirmRequestDto) ([]byte, error) {
	params := &stripe.PaymentIntentConfirmParams{
		PaymentMethodData: &stripe.PaymentIntentPaymentMethodDataParams{
			Card: &stripe.PaymentMethodCardParams{
				Token:    stripe.String(requestDto.CardToken),
			},
			BillingDetails: &stripe.BillingDetailsParams{
				Name: stripe.String(requestDto.CardName),
			},
			Type: stripe.String(string((stripe.PaymentMethodTypeCard))),
		},
	}

	pic, err := s.DoPaymentIntent.Confirm(requestDto.PaymentId, params)

	if err != nil {
		return []byte(""), err
	}

	paymentintentJson, err := json.Marshal(pic)

	if err != nil  {
		return []byte(""), err
	}

	return paymentintentJson, err
}

func (s Stripe) CardTokenGenerate(requestDto dto.CardTokenGenerateRequestDto) (string, error) {
	params := &stripe.TokenParams{
		Card: &stripe.CardParams{
			Number: stripe.String(requestDto.CardNumber),
			ExpMonth: stripe.String(requestDto.CardExpiryMonth),
			ExpYear: stripe.String(requestDto.CardExpiryYear),
			CVC: stripe.String(requestDto.CardSecurityCode),
		},
	}

	t, err := s.Token.New(params)

	if err != nil {
		return "", err
	}

	return t.ID, err
}
