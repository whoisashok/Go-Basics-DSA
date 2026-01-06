package main

import (
	"fmt"
	"sync"
)

// singleton represents the single instance of our object.
type singleton struct {
	data string
}

var (
	instance *singleton
	once     sync.Once
)

// GetInstance returns the singleton instance.
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{data: "Initial Data"}
		fmt.Println("Singleton instance created.")
	})
	return instance
}

func Singleton() {
	// Get the instance multiple times
	s1 := GetInstance()
	fmt.Printf("Instance 1 data: %s\n", s1.data)

	s2 := GetInstance()
	fmt.Printf("Instance 2 data: %s\n", s2.data)

	// Verify that both instances are the same
	if s1 == s2 {
		fmt.Println("Both instances are the same.")
	} else {
		fmt.Println("Instances are different.")
	}

	// Modify data through one instance, observe changes in the other
	s1.data = "Modified Data"
	fmt.Printf("Instance 2 data after modification: %s\n", s2.data)
}
