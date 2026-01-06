package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/sony/gobreaker"
)

// Define a global circuit breaker instance
var cb *gobreaker.CircuitBreaker

func init() {
	// Configure the circuit breaker settings
	settings := gobreaker.Settings{
		Name: "HTTPBreaker",
		// Timeout is the period of the open state, after which the state becomes half-open.
		Timeout: 5 * time.Second,
		// ReadyToTrip is called with a copy of Counts whenever a request fails in the closed state.
		// The default function trips the breaker if 5 consecutive failures occur.
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures > 5
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			fmt.Printf("Circuit Breaker '%s' state changed from %s to %s\n", name, from, to)
		},
	}
	cb = gobreaker.NewCircuitBreaker(settings)
}

// unreliableServiceCall simulates an external service call that might fail
func unreliableServiceCall(url string) ([]byte, error) {
	// Wrap the service call in the circuit breaker's Execute method.
	body, err := cb.Execute(func() (interface{}, error) {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode >= http.StatusInternalServerError {
			return nil, errors.New("service returned a server error")
		}

		return io.ReadAll(resp.Body)
	})

	if err != nil {
		return nil, err
	}

	// Type assertion for the result from Execute()
	return body.([]byte), nil
}

func CircuitBreaker() {
	// This example assumes a local server running on port 8080.
	// For demonstration, you would simulate failures on the server side.
	targetURL := "http://localhost:8080/data"

	for i := 0; i < 20; i++ {
		fmt.Printf("Request %d: ", i+1)
		data, err := unreliableServiceCall(targetURL)
		if err != nil {
			log.Printf("Error: %v. Breaker State: %s\n", err, cb.State())
		} else {
			fmt.Printf("Success. Data size: %d bytes. Breaker State: %s\n", len(data), cb.State())
		}
		time.Sleep(1 * time.Second)
	}
}
