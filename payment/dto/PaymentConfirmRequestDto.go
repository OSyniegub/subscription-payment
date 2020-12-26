package dto

type PaymentConfirmRequestDto struct {
	PaymentId        string `validate:"required"`
	Currency         string `validate:"required"`
	CardName         string `string:"card_name"`
	CardToken        string `validate:"required"`
}
