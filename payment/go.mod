module github.com/OSyniegub/subscription-payment/payment

go 1.15

require (
    github.com/OSyniegub/subscription-payment/payment/dto v1.0.0
	github.com/stripe/stripe-go/v71 v71.48.0
)

replace (
   github.com/OSyniegub/subscription-payment/payment/dto => ./dto
)