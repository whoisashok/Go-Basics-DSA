package main

import "fmt"

// SpiralMatrix generates an n x n matrix filled in a spiral order.
func SpiralMatrixSetNumbars(n int) [][]int {
	// Initialize an n x n matrix with zeros.
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	// Define the boundaries of the matrix.
	top := 0
	bottom := n - 1
	left := 0
	right := n - 1

	value := 1

	for value <= n*n {
		// Traverse top row (left to right)
		for i := left; i <= right; i++ {
			matrix[top][i] = value
			value++
		}
		top++ // Move the top boundary down

		// Traverse right column (top to bottom)
		for i := top; i <= bottom; i++ {
			matrix[i][right] = value
			value++
		}
		right-- // Move the right boundary left

		// Traverse bottom row (right to left)
		for i := right; i >= left; i-- {
			matrix[bottom][i] = value
			value++
		}
		bottom-- // Move the bottom boundary up

		// Traverse left column (bottom to top)
		for i := bottom; i >= top; i-- {
			matrix[i][left] = value
			value++
		}
		left++ // Move the left boundary right
	}
	return matrix
}

func SpiralMatrixDemo() {
	size := 4 // Example size
	spiral := SpiralMatrixSetNumbars(size)

	// Print the generated spiral matrix
	for i := 0; i < size; i++ {
		fmt.Println(spiral[i])
	}
}
