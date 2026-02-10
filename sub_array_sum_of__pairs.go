package main

import (
	"fmt"
	"sort"
)

func findPairsWithSumHashMap(arr []int, targetSum int) [][]int {
	seen := make(map[int]int) // Maps element value to its index
	pairs := [][]int{}        // To store pairs

	for _, num := range arr {
		complement := targetSum - num
		if _, found := seen[complement]; found {
			// Found a pair and append to return
			pairs = append(pairs, []int{num, complement})
			// Decrement count or remove to avoid reusing the same element
			seen[complement]--
		}
		// Add current number to the map
		seen[num]++
	}
	return pairs
}

func findPairsWithSumTwoPointers(arr []int, targetSum int) [][]int {
	pairs := [][]int{}
	sort.Ints(arr) // Sort the array

	left := 0
	right := len(arr) - 1

	for left < right {
		currentSum := arr[left] + arr[right]
		if currentSum == targetSum {
			pairs = append(pairs, []int{arr[left], arr[right]})
			left++
			right--
		} else if currentSum < targetSum {
			left++
		} else { // currentSum > targetSum
			right--
		}
	}
	return pairs
}

func minSubArrayLen(nums []int, target int) {
	left := 0
	sum := 0
	minLen := len(nums) + 1

	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		// Shrink window while sum >= target
		for sum >= target {
			if right-left+1 < minLen {
				minLen = right - left + 1
			}
			// Remove leftmost element from sum and move left pointer
			sum -= nums[left]
			left++
		}
	}

	if minLen == len(nums)+1 {
		fmt.Println("No valid subarray found.")
	}
	fmt.Printf("Minimum length of subarray: %d\n", minLen)
}

func ArraySumPairs() {
	arr1 := []int{4, 3, 6, 7, 8, 1, 9}
	target1 := 15
	fmt.Printf("Array: %v, Target: %d\n", arr1, target1)
	result1 := findPairsWithSumHashMap(arr1, target1)
	fmt.Println(result1)

	arr2 := []int{4, 3, 6, 7, 8, 1, 9}
	target2 := 15
	fmt.Printf("Array: %v, Target: %d\n", arr2, target2)
	result2 := findPairsWithSumTwoPointers(arr2, target2)
	fmt.Println(result2)

	arr3 := []int{1, 4, -4, 0, 2, 2}
	target3 := 4
	fmt.Printf("Array: %v, Target: %d\n", arr3, target3)
	minSubArrayLen(arr3, target3)
}
