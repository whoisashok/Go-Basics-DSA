package main

import "fmt"

func fibonacciIterative(n int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a)
		a, b = b, a+b // Update a and b for the next iteration
	}
	fmt.Println()
}

func Fibonacci() {
	fmt.Println("Fibonacci series up to 10 terms:")
	fibonacciIterative(10) // Prints the first 10 Fibonacci numbers
}
