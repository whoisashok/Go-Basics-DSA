package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type State string

const (
	StateClosed   State = "CLOSED"    // Everything is fine
	StateOpen     State = "OPEN"      // Service down, requests blocked
	StateHalfOpen State = "HALF-OPEN" // Testing if service is back
)

type CircuitBreaker struct {
	mutex            sync.Mutex
	state            State
	failureThreshold int
	failureCount     int
	lastFailureTime  time.Time
	retryTimeout     time.Duration
}

func NewCircuitBreaker(threshold int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:            StateClosed,
		failureThreshold: threshold,
		retryTimeout:     timeout,
	}
}

func (cb *CircuitBreaker) Execute(action func() error) error {
	cb.mutex.Lock()

	// Check if we should move from OPEN to HALF-OPEN
	if cb.state == StateOpen && time.Since(cb.lastFailureTime) > cb.retryTimeout {
		cb.state = StateHalfOpen
	}

	if cb.state == StateOpen {
		cb.mutex.Unlock()
		return errors.New("circuit is open: request blocked")
	}
	cb.mutex.Unlock()

	// Try the actual work
	err := action()

	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	if err != nil {
		cb.failureCount++
		cb.lastFailureTime = time.Now()
		if cb.failureCount >= cb.failureThreshold {
			cb.state = StateOpen
		}
		return err
	}

	// Success! Reset the breaker
	cb.state = StateClosed
	cb.failureCount = 0
	return nil
}

func CircuitBreakerImpl() {
	// Allow 3 failures before tripping; retry after 5 seconds
	cb := NewCircuitBreaker(3, 5*time.Second)

	// Mock function that always fails
	failingService := func() error {
		return errors.New("service timeout")
	}

	for i := 1; i <= 5; i++ {
		err := cb.Execute(failingService)
		fmt.Printf("Attempt %d: %v | Current State: %s\n", i, err, cb.state)
	}
}
