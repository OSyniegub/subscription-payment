package main

import (
	"encoding/json"
	"gateway/dto"
	"github.com/gorilla/mux"
	"github.com/stripe/stripe-go/v71/paymentintent"
	"io/ioutil"
	"net/http"
	"html/template"
	"github.com/stripe/stripe-go/v71"
)

type PaymentGateway interface {
	PaymentIntent(itemId string) (string, error)
	PaymentConfirm(clientSecret string) (stripe.PaymentIntent, error)
}

type Stripe struct {}
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

func (s Stripe) PaymentIntent(itemId string) (string, error) {
	stripe.Key = "sk_test_51I1dzaAoiWfQjN7OE4ExtBtv6S5RvXxcQQt8sIHzcMSfs9wgUakNFl5udNXckUHXvcLeVWY1wMzdAsfkJnhm5WQI00pOFESNLQ"

	params := &stripe.PaymentIntentParams{
		Amount: stripe.Int64(getAmount(itemId)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
	}

	params.AddMetadata("integration_check", "accept_a_payment")

	pi, err := paymentintent.New(params)

	return pi.ClientSecret, err
}

func (s Stripe) PaymentConfirm(clientSecret string) (stripe.PaymentIntent, error) {
	stripe.Key = "sk_test_51I1dzaAoiWfQjN7OE4ExtBtv6S5RvXxcQQt8sIHzcMSfs9wgUakNFl5udNXckUHXvcLeVWY1wMzdAsfkJnhm5WQI00pOFESNLQ"

	params := &stripe.PaymentIntentConfirmParams{
		PaymentMethodData: &stripe.PaymentIntentPaymentMethodDataParams{
			Card: &stripe.PaymentMethodCardParams{

			},
			BillingDetails: &stripe.BillingDetailsParams{

			},
		},
	}

	pic, err := paymentintent.Confirm(clientSecret, params)

	return *pic, err
}

func setupRouter() *mux.Router {
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/payment_form", paymentForm).Methods("GET")
	router.HandleFunc("/api/v1/charge", charge).Methods("POST")

	return router
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, _  := template.ParseFiles("assets/templates/home.html")
	tmpl.Execute(w, nil)
	return
}

func makePaymentIntent(pg PaymentGateway, itemId string) (string, error)  {
	return pg.PaymentIntent(itemId)
}

func makePaymentConfirm(pg PaymentGateway, clientSecret string) (stripe.PaymentIntent, error)  {
	return pg.PaymentConfirm(clientSecret)
}

func paymentForm(w http.ResponseWriter, r *http.Request) {
	itemId := r.URL.Query().Get("item_id")

	clientSecret, err := makePaymentIntent(&Stripe{}, itemId)

	if err != nil {
		//
	}

	paymentConfirm, err := makePaymentConfirm(&Stripe{}, clientSecret)
	
	if err != nil {
		//
	}

	tmpl, _  := template.ParseFiles("assets/templates/payment_form.html")
	tmpl.Execute(w, paymentConfirm)
	return
}

func charge(w http.ResponseWriter, r *http.Request) {
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	json.NewEncoder(w).Encode(b)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var authorizationRequestDto dto.PaymentRequestDto
	err = json.Unmarshal(b, &authorizationRequestDto)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(authorizationRequestDto)

	return
}
