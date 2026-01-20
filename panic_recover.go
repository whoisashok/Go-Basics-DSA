package main

import "fmt"

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r) // Catches the panic
	}
}

func mayPanic() {
	fmt.Println("Inside mayPanic")
	panic("a problem in mayPanic") // Panic occurs
}

// main function to demonstrate panic and recover
func PanicRecover() {
	defer handlePanic() // Defer the panic handler
	defer fmt.Println("Calling a function that might panic...")
	mayPanic()
	fmt.Println("This line will execute after recovery.") // Execution continues here
}
