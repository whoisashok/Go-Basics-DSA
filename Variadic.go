package main

import "fmt"

// Variadic function to calculate sum
func Sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func variadic() {
	fmt.Println("Sum of 1, 2, 3:", Sum(1, 2, 3))
	fmt.Println("Sum of 4, 5:", Sum(4, 5))
	fmt.Println("Sum of no numbers:", Sum())
}
