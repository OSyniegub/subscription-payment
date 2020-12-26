package payment

import (
	"errors"
	"github.com/OSyniegub/subscription-payment/payment/dto"
)

type Paypal struct {}

	func (s Paypal) PaymentIntent(amount string) (string, error) {
		/* TODO fix authorization error
			c, err := paypal.NewClient("Ae-XFaXQmO6HP40gB0CPXpYMDABYA5pfIca-E9qQng2_maYbiHmLP-QNtTLoQugG4nXExuV5hff-cFq_", "EGLkNyQrclpJJVNohNyXWy9j8nwkmbK0BacUN_nxk7ARoBuq1lTM9i_4bU2Zvj5BG_5D5u5OraRO2RMG", paypal.APIBaseSandBox)

			if err != nil {
				return "", err
			}

			order, err := c.CreateOrder(paypal.OrderIntentCapture, []paypal.PurchaseUnitRequest{
					{
						Amount: &paypal.PurchaseUnitAmount{
							Currency:  "USD",
							Value:     string(amount),
						},
					},
				},
				&paypal.CreateOrderPayer{
					Name:         nil,
					EmailAddress: "",
					PayerID:      "",
					Phone:        nil,
					BirthDate:    "",
					TaxInfo:      nil,
					Address:      nil,
				},
				&paypal.ApplicationContext{
					BrandName:          "",
					Locale:             "",
					LandingPage:        "",
					ShippingPreference: "",
					UserAction:         "",
					ReturnURL:          "",
					CancelURL:          "",
				})

			if err != nil {
				fmt.Println(err)
				return "", err
			}

			return order
		*/
		return "", errors.New("")
	}

func (s Paypal) PaymentConfirm(requestDto dto.PaymentConfirmRequestDto) ([]byte, error) {
	//TODO implement PayPal payment processing
	//For now it uses Stripe (hardcoded)
	return []byte(""), errors.New("")
}

func (s Paypal) CardTokenGenerate(requestDto dto.CardTokenGenerateRequestDto) (string, error) {
	//TODO implement PayPal payment processing
	//For now it uses Stripe (hardcoded)
	return "", errors.New("")
}