package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func routes() {
	r := mux.NewRouter()

	// Histogram Endpoint
	r.HandleFunc("CSV/histogram", HistogramHandler).Methods("GET")

	// Passenger Data Endpoints
	r.HandleFunc("CSV/passengers", CSVPassengerHandler).Methods("GET")

	// SQLite Passenger Data Endpoints
	r.HandleFunc("sqlite/passenger", GetPassengers).Methods("GET")
	r.HandleFunc("sqlite/passenger/{passengerId}", GetPassengerById).Methods("GET")

	http.Handle("/", r)

	// Start the server on port 8080
	fmt.Println("Server listening on :8080...")
	http.ListenAndServe(":8080", nil)
}
