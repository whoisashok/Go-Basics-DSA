package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// --- Structs for Request and Response ---

// Define the struct for the data we send (both Server and Client use this)
type RequestPayload struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

// Define the struct for the response data (both Server and Client use this)
type ResponsePayload struct {
	Status    string         `json:"status"`
	Received  RequestPayload `json:"received_data"`
	Timestamp string         `json:"timestamp"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Handling a GET request\n")
		// code to handle GET request
	case "POST":
		fmt.Fprintf(w, "Handling a POST request\n")
		// code to handle POST request
	case "PUT":
		fmt.Fprintf(w, "Handling a PUT request\n")
		// code to handle PUT request
	case "DELETE":
		fmt.Fprintf(w, "Handling a DELETE request\n")
		// code to handle DELETE request
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2. Decode the incoming JSON payload
	var payload RequestPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid request body format (must be JSON)", http.StatusBadRequest)
		return
	}

	// 3. Log the received data
	fmt.Printf("\n✅ [SERVER] Received Data: Key='%s', Value=%d\n", payload.Key, payload.Value)

	// 4. Prepare the response
	response := ResponsePayload{
		Status:    "Success: Data processed by server",
		Received:  payload,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	// 5. Set the response header and status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // HTTP 201 Created

	// 6. Encode and send the JSON response
	json.NewEncoder(w).Encode(response)
}

func main() {

	// Register the handler function for the /submit path
	http.HandleFunc("/", handler)
	http.HandleFunc("/submit", submitHandler)

	port := ":8080"
	fmt.Printf("🚀 [SERVER] Starting on http://localhost%s\n", port)

	// Start the server (this is a BLOCKING call)
	// log.Fatal ensures the program terminates gracefully on error or when explicitly stopped by runClient
	log.Fatal(http.ListenAndServe(port, nil))
}
