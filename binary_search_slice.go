package main

import (
	"fmt"
	"slices"
)

func binarySearchSlice() {
	sortedArray1 := []string{"apple", "banana", "cherry", "date", "elderberry"}
	target1 := "cherry"

	// slices.BinarySearch returns the index if found, or the insertion point if not found.
	// It also returns a boolean indicating if the value was found.
	index1, found1 := slices.BinarySearch(sortedArray1, target1)
	if found1 {
		fmt.Printf("Target \"%s\" found at index %d\n", target1, index1)
	} else {
		fmt.Printf("Target \"%s\" not found (would be inserted at index %d)\n", target1, index1)
	}

	sortedArray2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target2 := 3

	// slices.BinarySearch returns the index if found, or the insertion point if not found.
	// It also returns a boolean indicating if the value was found.
	index2, found2 := slices.BinarySearch(sortedArray2, target2)
	if found2 {
		fmt.Printf("Target \"%d\" found at index %d\n", target2, index2)
	} else {
		fmt.Printf("Target \"%d\" not found (would be inserted at index %d)\n", target2, index2)
	}
}
