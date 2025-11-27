package main

import (
	"fmt"
	"sort"
)

func findPairsWithSumHashMap(arr []int, targetSum int) {
	seen := make(map[int]int) // Maps element value to its index

	for i, num := range arr {
		complement := targetSum - num
		if _, found := seen[complement]; found {
			fmt.Printf("Pair found: (%d, %d) at indices (%d, %d)\n", complement, num, seen[complement], i)
			return // Return after finding the first pair, or remove 'return' to find all pairs
		}
		seen[num] = i
	}
	fmt.Println("No pair found with the given sum.")
}

func findPairsWithSumTwoPointers(arr []int, targetSum int) {
	sort.Ints(arr) // Sort the array

	left := 0
	right := len(arr) - 1

	for left < right {
		currentSum := arr[left] + arr[right]

		if currentSum == targetSum {
			fmt.Printf("Pair found: (%d, %d)\n", arr[left], arr[right])
			return // Return after finding the first pair, or remove 'return' to find all pairs
		} else if currentSum < targetSum {
			left++
		} else { // currentSum > targetSum
			right--
		}
	}
	fmt.Println("No pair found with the given sum.")
}

func ArraySumPairs() {
	arr1 := []int{4, 3, 6, 7, 8, 1, 9}
	target1 := 15
	fmt.Printf("Array: %v, Target: %d\n", arr1, target1)
	findPairsWithSumHashMap(arr1, target1)

	arr2 := []int{1, 2, 3, 4, 5}
	target2 := 10
	fmt.Printf("\nArray: %v, Target: %d\n", arr2, target2)
	findPairsWithSumHashMap(arr2, target2)

	arr3 := []int{4, 3, 6, 7, 8, 1, 9}
	target3 := 15
	fmt.Printf("Array: %v, Target: %d\n", arr3, target3)
	findPairsWithSumTwoPointers(arr3, target3)

	arr4 := []int{1, 2, 3, 4, 5}
	target4 := 10
	fmt.Printf("\nArray: %v, Target: %d\n", arr4, target4)
	findPairsWithSumTwoPointers(arr4, target4)
}
