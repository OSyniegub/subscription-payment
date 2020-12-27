module github.com/OSyniegub/subscription-payment

go 1.15

require (
	github.com/OSyniegub/subscription-payment/payment v1.0.0
	github.com/OSyniegub/subscription-payment/payment/dto v1.0.0
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/google/uuid v1.1.2
	github.com/gorilla/csrf v1.7.0 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/logpacker/PayPal-Go-SDK v2.0.5+incompatible // indirect
	github.com/stripe/stripe-go/v71 v71.48.0
	golang.org/x/sys v0.0.0-20201223074533-0d417f636930 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0
)

replace (
	github.com/OSyniegub/subscription-payment/payment => ./payment
	github.com/OSyniegub/subscription-payment/payment/dto => ./payment/dto
)
