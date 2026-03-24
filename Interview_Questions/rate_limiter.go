package main

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

// Simple rate limiter using time.Tick
func RateLimiter() {
	limiter := time.Tick(500 * time.Millisecond)

	for i := 0; i < 5; i++ {
		<-limiter
		println("request", i)
	}
}

// More robust rate limiter using golang.org/x/time/rate
var limiter = rate.NewLimiter(2, 4) // 2 req/sec

func rateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
