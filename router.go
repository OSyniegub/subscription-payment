package main

import (
	"crypto/md5"
	"fmt"
	"github.com/OSyniegub/subscription-payment/payment"
	"github.com/OSyniegub/subscription-payment/payment/dto"
	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
	"html/template"
	"io"
	"net/http"
	"strconv"
)

func setupRouter() *mux.Router {
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/payment_form", paymentForm).Methods("GET")
	// API v1
	router.HandleFunc("/api/v1/charge", paymentCharge).Methods("POST")

	return router
}

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

func getPaymentGateway(paymentMethod string) payment.Gateway {
	/* not working due to authorization error
	var gateway payment.Gateway

	if paymentMethod == "card" {
		gateway = &payment.Stripe{}
	} else if paymentMethod == "paypal" {
			gateway = &payment.Paypal{}
	}
	*/

	return &payment.Stripe{}
}

var token string

func generateToken() string {
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(3, 10))
	io.WriteString(h, "tok")
	token := fmt.Sprintf("%x", h.Sum(nil))

	return token
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, _  := template.ParseFiles("assets/templates/home.html")
	tmpl.Execute(w, nil)
	return
}

func paymentForm(w http.ResponseWriter, r *http.Request) {
	itemId := r.URL.Query().Get("item_id")

	token = generateToken()

	tmpl, _  := template.ParseFiles("assets/templates/payment_form.html")
	tmpl.Execute(w, map[string]interface{}{
		"Amount": getAmount(itemId),
		"Token": token,
	})
	return
}

func paymentCharge(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Form.Get("csrf") != token {
		http.Error(w, "Forbidden csrf token", 500)
		return
	}

	gateway := getPaymentGateway(r.Form.Get("payment_method"))

	paymentId, err := payment.MakePaymentIntent(gateway, r.Form.Get("amount"))

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	cardToken := r.Form.Get("card_token")

	if cardToken == "" {
		cardTokenGenerateRequestDto := dto.CardTokenGenerateRequestDto{
			CardNumber:       r.Form.Get("card_number"),
			CardExpiryMonth:  r.Form.Get("card_expiry_month"),
			CardExpiryYear:   r.Form.Get("card_expiry_year"),
			CardSecurityCode: r.Form.Get("card_security_code"),
		}

		err := validator.New().Struct(cardTokenGenerateRequestDto)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		cardToken, err = payment.MakeCardTokenGenerate(&payment.Stripe{}, cardTokenGenerateRequestDto)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}

	// As far as JS script receive card token from google pay test env we need to replace it with stripe test card token
	// Will be used only in case of GooglePay method
	if cardToken == "examplePaymentMethodToken" {
		cardToken = "tok_visa"
	}

	paymentConfirmRequestDto := dto.PaymentConfirmRequestDto{
		PaymentId: paymentId,
		Currency:  r.Form.Get("currency"),
		CardName:  r.Form.Get("card_name"),
		CardToken: cardToken,
	}

	err = validator.New().Struct(paymentConfirmRequestDto)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	paymentConfirm, err := payment.MakePaymentConfirm(&payment.Stripe{}, paymentConfirmRequestDto)

	if err != nil  {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(paymentConfirm)
}