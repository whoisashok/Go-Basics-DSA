package main

import "time"

// ProcessNumber processes a number and returns the result through a channel
func ProcessNumber(input int, output chan<- int) {
	go func() {
		// Simulate a time-consuming task
		time.Sleep(2 * time.Second)
		output <- input * 2
	}()
}

func main() {
	// Example usage
	output := make(chan int)
	ProcessNumber(5, output)
	result := <-output
	println("Processed result:", result)
}
