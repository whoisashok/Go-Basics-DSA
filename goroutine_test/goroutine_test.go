package main

import (
	"testing"
	"time"
)

func TestProcessNumber(t *testing.T) {
	expected := 10
	input := 5
	output := make(chan int)

	ProcessNumber(input, output)

	select {
	case result := <-output:
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("Test timed out")
	}
}

// Best Practices for Testing Goroutines
// =====================================
// Timeouts:
// 	Always implement timeouts using the select statement to avoid tests that hang indefinitely.

// Synchronization:
// 	Use channels or other synchronization techniques like sync.WaitGroup to coordinate between your main testing goroutine and the child goroutines.

// Deterministic Inputs and Outputs:
// 	Ensure your tests are deterministic by carefully selecting inputs and controlling the environment.
