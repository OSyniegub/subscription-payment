package main

import (
	"github.com/OSyniegub/subscription-payment/gateway"
	"github.com/gorilla/mux"
	"github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/charge"
	"github.com/stripe/stripe-go/v71/paymentintent"
	"html/template"
	"net/http"
)

type ChargeRequestDto struct {
	PaymentId string `json:"payment_id"`
}

type ChargeResponseDto struct {
	Status string `json:"status"`
	ReceiptUrl string `njson:"charges.data.receipt_url"`
}

type StripeChargeDataDto struct {
	ReceiptUrl string `json:"receipt_url"`
}

type ApplePay struct {}
type GooglePay struct {}
type PayPal struct {}

func getAmount(itemId string) int64 {
	var amount int64

	if itemId == "1" {
		amount = 399
	} else if itemId == "2" {
		amount = 499
	} else {
		amount = 599
	}

	return amount
}

func setupRouter() *mux.Router {
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/payment_form", paymentForm).Methods("GET")
	router.HandleFunc("/api/v1/charge", paymentCharge).Methods("POST")

	return router
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, _  := template.ParseFiles("assets/templates/home.html")
	tmpl.Execute(w, nil)
	return
}

func makePaymentIntent(pg PaymentGateway, itemId string) (string, error)  {
	return pg.PaymentIntent(getAmount(itemId))
}

func makePaymentConfirm(pg PaymentGateway, paymentId string) (stripe.PaymentIntent, error)  {
	return pg.PaymentConfirm(paymentId)
}

func getCharge(pg PaymentGateway, chargeId string) (string, error)  {
	return pg.ChargeGet(chargeId)
}

func paymentForm(w http.ResponseWriter, r *http.Request) {
	itemId := r.URL.Query().Get("item_id")

	paymentId, err := makePaymentIntent(&Stripe{}, itemId)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tmpl, _  := template.ParseFiles("assets/templates/payment_form.html")
	tmpl.Execute(w, paymentId)
	return
}

func paymentCharge(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	paymentId := r.Form.Get("payment_id")

	paymentConfirm, err := makePaymentConfirm(&Stripe{}, paymentId)

	if err != nil && paymentConfirm.Status != "succeed" {
		http.Error(w, err.Error(), 500)
		return
	}

	http.Redirect(w, r, paymentConfirm.Charges.Data[0].ReceiptURL, http.StatusSeeOther)

	return
}