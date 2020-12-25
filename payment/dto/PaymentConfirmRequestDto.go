package dto

type PaymentConfirmRequestDto struct {
	PaymentId        string `validate:"required"`
	Currency         string `validate:"required"`
	CardName         string `validate:"required"`
	CardToken        string `validate:"required"`
}
