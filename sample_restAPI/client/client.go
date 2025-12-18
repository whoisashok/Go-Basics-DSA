package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

// --- Client Logic ---

func runClient() {
	// Give the server a moment to start listening
	time.Sleep(1 * time.Second)

	fmt.Println("\n--- 🚀 [CLIENT] Starting POST Request ---")

	// 1. Prepare the data to send
	dataToSend := RequestPayload{
		Key:   "TransactionID-999",
		Value: 42,
	}

	// Marshal the Go struct into a JSON byte slice
	jsonData, err := json.Marshal(dataToSend)
	if err != nil {
		fmt.Printf("❌ [CLIENT] Error marshalling data: %v\n", err)
		return
	}

	// The target API endpoint
	apiURL := "http://localhost:8080/submit"

	// 2. Create the POST request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("❌ [CLIENT] Error creating request: %v\n", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		// This is the common error if the server didn't start in time
		fmt.Printf("❌ [CLIENT] Error sending request (Did the server start?): %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 3. Check the response status
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		fmt.Printf("❌ [CLIENT] API request failed with status code: %d\n", resp.StatusCode)
		return
	}

	// 4. Decode the JSON response
	var responseData ResponsePayload
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		fmt.Printf("❌ [CLIENT] Error decoding response: %v\n", err)
		return
	}

	// 5. Print the results
	fmt.Println("\n--- ✅ [CLIENT] Response Received ---")
	fmt.Printf("Status: %s\n", responseData.Status)
	fmt.Printf("Server Time: %s\n", responseData.Timestamp)
	fmt.Println("\nServer Echoed Data:")
	fmt.Printf("Key: %s\n", responseData.Received.Key)
	fmt.Printf("Value: %d\n", responseData.Received.Value)

	// Since the server will block forever, we exit the whole program after the client runs
	// This is only for a single-file demo; in production, the server would run indefinitely.
	fmt.Println("\n--- Demo Complete. Exiting Server ---")

	// This will cause the log.Fatal below to exit the main goroutine (and the program)
	// You wouldn't do this in a real server application.
	// You need a way to stop the infinite loop of the server, hence the log.Fatal(errors.New)
	log.Fatal("Client finished, shutting down demo server.")
}

func simpleClient() {
	jsonData := []byte(`{"key": "value"}`)
	bodyReader := bytes.NewReader(jsonData)

	response, err := http.Post("http://localhost:8080/", "application/json", bodyReader)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Status Code: %d\n", response.StatusCode)
	fmt.Printf("Response Body: %s\n", body)
}

// --- Main Execution ---
func main() {

	// Start the client in a separate goroutine
	// This allows the main goroutine to proceed to start the server.
	go runClient()
	time.Sleep(100 * time.Millisecond) // Small delay to ensure client starts after server

	// Start the simple client
	simpleClient()
}
