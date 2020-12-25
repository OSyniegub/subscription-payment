package gateway

import (
	"github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/charge"
	"github.com/stripe/stripe-go/v71/paymentintent"
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

	return pi.ID, err
}

func (s Stripe) PaymentConfirm(paymentId string) (string, error) {
	stripe.Key = "sk_test_51I1dzaAoiWfQjN7OE4ExtBtv6S5RvXxcQQt8sIHzcMSfs9wgUakNFl5udNXckUHXvcLeVWY1wMzdAsfkJnhm5WQI00pOFESNLQ"

	/*
		TODO add logic to generate card token and remove commented code below
		cardCVC := "323"
		cardExpMonth := "12"
		cardExpYear := "2021"
		cardNumber := "4242424242424242"
	*/
	cardToken := "tok_visa"

	params := &stripe.PaymentIntentConfirmParams{
		PaymentMethodData: &stripe.PaymentIntentPaymentMethodDataParams{
			Card: &stripe.PaymentMethodCardParams{
				/*
					CVC:      stripe.String(cardCVC),
					ExpMonth: stripe.String(cardExpMonth),
					ExpYear:  stripe.String(cardExpYear),
					Number:   stripe.String(cardNumber),
				*/
				Token:    stripe.String(cardToken),
			},
			BillingDetails: &stripe.BillingDetailsParams{

			},
			Type: stripe.String("card"),
		},
	}

	pic, err := paymentintent.Confirm(paymentId, params)

	return pic.Charges.Data[0].ReceiptURL, err
}

func (s Stripe) ChargeGet(chargeId string) (string, error) {
	stripe.Key = "sk_test_51I1dzaAoiWfQjN7OE4ExtBtv6S5RvXxcQQt8sIHzcMSfs9wgUakNFl5udNXckUHXvcLeVWY1wMzdAsfkJnhm5WQI00pOFESNLQ"

	c, err := charge.Get(chargeId, nil)

	return c.ReceiptURL, err
}
