package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/time/rate"
)

// Message is a simple struct for JSON responses
type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

// Global rate limiter instance
var limiter = rate.NewLimiter(2, 4) // Allows 2 events per second, max burst size of 4

// limitMiddleware is an http.Handler middleware that rate limits requests
func limitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if a token can be taken from the bucket immediately
		if !limiter.Allow() {
			w.WriteHeader(http.StatusTooManyRequests) // 429 status code
			fmt.Fprintf(w, `{"status": "Request Failed", "body": "The API is at capacity, try again later."}`)
			return
		}

		// If allowed, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}

// helloHandler is the actual handler for the "/" endpoint
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func RateLimiter() {
	mux := http.NewServeMux()
	// Wrap the helloHandler with the rate limiting middleware
	mux.Handle("/", limitMiddleware(http.HandlerFunc(helloHandler)))

	log.Println("Listening on port 3000")
	// Start the server
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal(err)
	}
}
