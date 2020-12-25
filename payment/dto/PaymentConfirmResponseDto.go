package dto

type PaymentConfirmResponseDto struct {
	Status string `json:"status"`
	ReceiptUrl string `njson:"charges.data.receipt_url"`
}
