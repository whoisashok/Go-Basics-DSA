package main

import (
	"fmt"
	"sort"
)

func sortMatrix(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])

	// Flatten into one slice
	flat := make([]int, 0, rows*cols)
	for _, row := range matrix {
		flat = append(flat, row...)
	}

	// Sort the flattened slice
	sort.Ints(flat)

	// Rebuild matrix in-place
	idx := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = flat[idx]
			idx++
		}
	}
}

func SortMatrix() {
	matrix := [][]int{
		{3, 1, 2},
		{6, 4, 5},
	}

	sortMatrix(matrix)
	fmt.Println(matrix)
}
