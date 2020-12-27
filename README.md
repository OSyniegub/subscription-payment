# subscription-payment

## Requirements

- [Golang v1.15](https://golang.org/doc/install)

## Steps

- run `go get -d ./...` to get all dependencies
- run `go build` to prebuild all go files
- run `go test payment/*.go` to check tests

## Start

- run `./subscription-payment` It will run the project

## Stop

When you're done work with the project just press `^C` on a terminal where the project started to turn it off.

## Test credentials

### Stripe

- CardNumber - 4242424242424242
- Name, Expiration, CVV - any values that looks like real

### PayPal

- Email - sb-syty04333597@personal.example.com
- Password - av}5E&Ly

### GooglePay

- personal card can be used (it will be replaced with test card token)

### ApplePay

- Sorry, it doesn't support (Apple developer account needed)
- But if you will run the project from Safari you will see unclickable apple pay button :)
