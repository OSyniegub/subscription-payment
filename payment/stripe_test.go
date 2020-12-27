package payment

import (
	"encoding/json"
	"github.com/OSyniegub/subscription-payment/payment/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stripe/stripe-go/v71"
	"testing"
)

type StripePaymentIntentMock struct {}
type StripeTokenMock struct {}

func (s StripePaymentIntentMock) New(params *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error) {
	return &stripe.PaymentIntent{
		ID: "payment_id",
	}, nil
}

func (s StripePaymentIntentMock) Confirm(id string, params *stripe.PaymentIntentConfirmParams) (*stripe.PaymentIntent, error) {
	return &stripe.PaymentIntent{
		Status: "paid",
	}, nil
}

func (s StripeTokenMock) New(params *stripe.TokenParams) (*stripe.Token, error) {
	return &stripe.Token{
		ID:          "card_token",
	}, nil
}

var stripeMock = &Stripe{
	DoPaymentIntent: StripePaymentIntentMock{},
	Token: StripeTokenMock{},
}

func TestStripe_PaymentIntentNew(t *testing.T) {
	shouldReceive := "payment_id"

	newPaymentIntentId, _  := stripeMock.PaymentIntent("1500")

	assert.Equal(t, shouldReceive, newPaymentIntentId, "PaymentIntent.New should return payment_id")
}

func TestStripe_PaymentIntentConfirm(t *testing.T) {
	shouldReceive, _ := json.Marshal(&stripe.PaymentIntent{
		Status: "paid",
	})

	confirmPaymentIntent, _  := stripeMock.PaymentConfirm(dto.PaymentConfirmRequestDto{
		PaymentId: "payment_id",
		Currency:  "USD",
		CardName:  "Oleksandr",
		CardToken: "token",
	})

	assert.Equal(t, shouldReceive, confirmPaymentIntent, "PaymentIntent.Confirm should return status: paid")
}

func TestStripe_TokenNew(t *testing.T) {
	shouldReceive := "card_token"

	newToken, _  := stripeMock.CardTokenGenerate(dto.CardTokenGenerateRequestDto{
		CardNumber:       "test_number",
		CardExpiryMonth:  "test_exp_month",
		CardExpiryYear:   "test_exp_year",
		CardSecurityCode: "test_security_code",
	})

	assert.Equal(t, shouldReceive, newToken, "Token.New should return card_token")
}
