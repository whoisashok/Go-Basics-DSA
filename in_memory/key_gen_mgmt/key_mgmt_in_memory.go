package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Key struct {
	ID      string    `json:"id"`
	Expires time.Time `json:"expires"`
}

var store = map[string]Key{}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /keys", GetKeys)
	mux.HandleFunc("POST /keys", CreateKey)
	mux.HandleFunc("GET /keys/{id}", GetKey)
	mux.HandleFunc("DELETE /keys/{id}", DeleteKey)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func GetKeys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	now := time.Now()
	var res []Key

	for _, k := range store {
		if k.Expires.After(now) {
			res = append(res, k)
		}
	}

	json.NewEncoder(w).Encode(res)
}

func GetKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")
	k, ok := store[id]

	if !ok || time.Now().After(k.Expires) {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(k)
}

func CreateKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body struct {
		Expires *time.Time `json:"expires"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil && err.Error() != "EOF" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	exp := time.Now().Add(time.Hour)
	if body.Expires != nil {
		exp = *body.Expires
	}

	k := Key{
		ID:      uuid.New().String(),
		Expires: exp,
	}

	store[k.ID] = k

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(k)
}

func DeleteKey(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if _, ok := store[id]; !ok {
		http.NotFound(w, r)
		return
	}

	delete(store, id)
	w.WriteHeader(http.StatusNoContent)
}

/*
// CURLs

// Create Key (Default 1 hour expiry)
curl -X POST http://localhost:8080/keys

// Create Key With Custom Expiry
curl -X POST http://localhost:8080/keys \
  -H "Content-Type: application/json" \
  -d '{"expires":"2026-12-31T23:59:00Z"}'

// Get All Active Keys
curl http://localhost:8080/keys

// Get Key By ID
curl http://localhost:8080/keys/d3a1f0c7-1234-4567-8910-abcdef123456

// Delete Key By ID
curl -X DELETE http://localhost:8080/keys/d3a1f0c7-1234-4567-8910-abcdef123456

*/
