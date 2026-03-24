package main

import "fmt"

// SpiralMatrix generates an n x n matrix filled in a spiral order.
func SpiralMatrixPrint() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	// Define the boundaries of the matrix.
	top := 0
	bottom := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	for range len(matrix) * len(matrix[0]) {
		// Traverse top row (left to right)
		for i := left; i <= right; i++ {
			fmt.Println(matrix[top][i])
		}
		top++ // Move the top boundary down

		// Traverse right column (top to bottom)
		for i := top; i <= bottom; i++ {
			fmt.Println(matrix[i][right])
		}
		right-- // Move the right boundary left

		// Traverse bottom row (right to left)
		for i := right; i >= left; i-- {
			fmt.Println(matrix[bottom][i])
		}
		bottom-- // Move the bottom boundary up

		// Traverse left column (bottom to top)
		for i := bottom; i >= top; i-- {
			fmt.Println(matrix[i][left])
		}
		left++ // Move the left boundary right
	}
}
