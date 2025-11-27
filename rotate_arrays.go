package main

import "fmt"

// rotateLeft rotates a slice to the left by 'k' positions.
func rotateLeft(arr []int, k int) []int {
	arrlen := len(arr)
	if arrlen == 0 || k == 0 {
		return arr // No rotation needed for empty array or zero rotations
	}

	// Normalize k to be within the bounds of the array length
	k = k % arrlen
	fmt.Println("k:", k)

	// Create a new slice by appending the elements after k to the elements before k
	rotatedArr := append(arr[k:], arr[:k]...)
	return rotatedArr
}

func RotateArraysLeft() {
	arr := []int{1, 2, 3, 4, 5}
	k := 2 // Rotate by 2 positions to the left

	fmt.Println("Original array:", arr)

	rotated := rotateLeft(arr, k)
	fmt.Println("Rotated array:", rotated) // Output: [3 4 5 1 2]

	arr2 := []int{10, 20, 30, 40}
	k2 := 1
	rotated2 := rotateLeft(arr2, k2)
	fmt.Println("Rotated array 2:", rotated2) // Output: [20 30 40 10]

	arr3 := []int{1, 2, 3}
	k3 := 5 // k greater than array length
	rotated3 := rotateLeft(arr3, k3)
	fmt.Println("Rotated array 3:", rotated3) // Output: [3 1 2] (5 % 3 = 2 rotations)
}
