package main

import (
	"log"
	"net/http"
)

func userService(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Service"))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request:", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func Middleware() {
	http.Handle("/users", loggingMiddleware(authMiddleware(http.HandlerFunc(userService))))
}
