package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"html/template"
)

func setupRouter() *mux.Router {
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/payment_form", home).Methods("GET")

	return router
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, _  := template.ParseFiles("assets/templates/home.html")
	tmpl.Execute(w, nil)
	return
}

func paymentForm(w http.ResponseWriter, r *http.Request) {
	indexStab := "Subscription plan"

	tmpl, _  := template.ParseFiles("assets/templates/home.html")
	tmpl.Execute(w, indexStab)
	return
}
