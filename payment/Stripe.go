package payment

import (
	"github.com/OSyniegub/subscription-payment/payment/dto"
	"github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/paymentintent"
	"github.com/stripe/stripe-go/v71/token"
)

type Stripe struct {}

func (s Stripe) PaymentIntent(amount int64) (string, error) {
	stripe.Key = "sk_test_51I1dzaAoiWfQjN7OE4ExtBtv6S5RvXxcQQt8sIHzcMSfs9wgUakNFl5udNXckUHXvcLeVWY1wMzdAsfkJnhm5WQI00pOFESNLQ"

	params := &stripe.PaymentIntentParams{
		Amount: stripe.Int64(amount),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
	}

	params.AddMetadata("integration_check", "accept_a_payment")

	pi, err := paymentintent.New(params)

	if err != nil {
		return "", err
	}

	return pi.ID, err
}

func (s Stripe) PaymentConfirm(requestDto dto.PaymentConfirmRequestDto) (dto.PaymentConfirmResponseDto, error) {
	stripe.Key = "sk_test_51I1dzaAoiWfQjN7OE4ExtBtv6S5RvXxcQQt8sIHzcMSfs9wgUakNFl5udNXckUHXvcLeVWY1wMzdAsfkJnhm5WQI00pOFESNLQ"

	params := &stripe.PaymentIntentConfirmParams{
		PaymentMethodData: &stripe.PaymentIntentPaymentMethodDataParams{
			Card: &stripe.PaymentMethodCardParams{
				Token:    stripe.String(requestDto.CardToken),
			},
			BillingDetails: &stripe.BillingDetailsParams{
				Name: stripe.String(requestDto.CardName),
			},
			Type: stripe.String("card"),
		},
	}

	pic, err := paymentintent.Confirm(requestDto.PaymentId, params)

	if err != nil {
		return dto.PaymentConfirmResponseDto{}, err
	}

	return dto.PaymentConfirmResponseDto{
		Status:		string(pic.Status),
		ReceiptUrl:	pic.Charges.Data[0].ReceiptURL,
	}, err
}

func (s Stripe) CardTokenGenerate(requestDto dto.CardTokenGenerateRequestDto) (string, error) {
	stripe.Key = "sk_test_51I1dzaAoiWfQjN7OE4ExtBtv6S5RvXxcQQt8sIHzcMSfs9wgUakNFl5udNXckUHXvcLeVWY1wMzdAsfkJnhm5WQI00pOFESNLQ"

	params := &stripe.TokenParams{
		Card: &stripe.CardParams{
			Number: stripe.String(requestDto.CardNumber),
			ExpMonth: stripe.String(requestDto.CardExpiryMonth),
			ExpYear: stripe.String(requestDto.CardExpiryYear),
			CVC: stripe.String(requestDto.CardSecurityCode),
		},
	}

	t, err := token.New(params)

	if err != nil {
		return "", err
	}

	return t.ID, err
}
