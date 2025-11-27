package main

import "fmt"

func maxSubarraySum(nums []int, k int) int {
	if len(nums) < k || k <= 0 {
		return 0 // Handle invalid input
	}

	currentSum := 0
	// Calculate the sum of the first window
	for i := 0; i < k; i++ {
		currentSum += nums[i]
	}

	maxSum := currentSum

	// Slide the window
	for i := k; i < len(nums); i++ {
		currentSum = currentSum + nums[i]   // Add the new element
		currentSum = currentSum - nums[i-k] // Remove the element leaving the window
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

func SlidingWindowSum() {
	arr := []int{1, 5, 4, 8, 7, 1, 9}
	k := 3
	result := maxSubarraySum(arr, k)
	fmt.Printf("Maximum subarray sum of length %d: %d\n", k, result) // Output: 19
}
