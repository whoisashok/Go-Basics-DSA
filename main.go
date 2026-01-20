package main

import (
	"fmt"
	"log"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {

	//	REST API
	//handleRequests()

	//	INTERFACE
	Interfaces()
	Interfaces2()

	//	LONGEST SUBSTRING WITHOUT REPEATING CHARACTERS
	LongestSubstringNoRepeat()

	// 	REVERSE SLICE
	reverse()

	//	CHANNELS - SEND & RECEIVE
	SendReceiveChannels()

	//	GOROUTINES - ODD & EVEN
	OddEvenGoroutine()
}
