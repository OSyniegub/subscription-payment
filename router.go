package main

import (
	"github.com/OSyniegub/subscription-payment/payment"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

type ChargeRequestDto struct {
	PaymentId string `json:"payment_id"`
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

func paymentForm(w http.ResponseWriter, r *http.Request) {
	itemId := r.URL.Query().Get("item_id")

	paymentId, err := payment.MakePaymentIntent(&payment.Stripe{}, getAmount(itemId))

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

	paymentConfirm, err := payment.MakePaymentConfirm(&payment.Stripe{}, paymentId)

	if err != nil || paymentConfirm.Status != "succeed" {
		http.Error(w, err.Error(), 500)
		return
	}

	http.Redirect(w, r, paymentConfirm.ReceiptUrl, http.StatusSeeOther)

	return
}