package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// create reverse proxy
func newProxy(target string) *httputil.ReverseProxy {
	url, _ := url.Parse(target)
	return httputil.NewSingleHostReverseProxy(url)
}

func ReverseProxyAPIGateway() {
	userService := newProxy("http://localhost:8081")
	orderService := newProxy("http://localhost:8082")

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		userService.ServeHTTP(w, r)
	})

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		orderService.ServeHTTP(w, r)
	})

	log.Println("API Gateway running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
