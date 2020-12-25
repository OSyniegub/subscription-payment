package dto

type CardTokenGenerateRequestDto struct {
	CardNumber       string `validate:"required"`
	CardExpiryMonth  string `validate:"required"`
	CardExpiryYear   string `validate:"required"`
	CardSecurityCode string `validate:"required"`
}
