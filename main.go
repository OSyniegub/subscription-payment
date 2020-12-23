package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	return port
}

func main() {
	// Routes Configuration
	router := setupRouter()

	// Server Configuration
	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + getPort(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
