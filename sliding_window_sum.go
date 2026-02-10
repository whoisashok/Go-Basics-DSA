package main

import (
	"errors"
	"fmt"
)

func maxSubarraySum(arr []int, k int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("array is empty")
	}

	// Variable-size window (Kadane’s Algorithm)
	if k == 0 {
		current := arr[0]
		maxSum := arr[0]
		// Iterate through the array, updating the current sum and max sum
		for i := 1; i < len(arr); i++ {
			if current < 0 {
				current = arr[i]
			} else {
				current += arr[i]
			}

			if current > maxSum {
				maxSum = current
			}
		}
		return maxSum, nil
	}

	// Fixed-size window (Sliding Window)
	if k < 0 || k > len(arr) {
		return 0, errors.New("invalid window size")
	}

	sum := 0
	// Calculate the sum of the first 'k' elements
	for i := range k {
		sum += arr[i]
	}
	maxSum := sum

	// Slide the window across the array
	for i := k; i < len(arr); i++ {
		sum += arr[i]
		// Remove the element that is sliding out of the window
		sum -= arr[i-k]

		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum, nil
}

func SlidingWindowSum() {
	arr := []int{1, -2, 3, 4, -1, 2}

	// Variable window
	maxVar, _ := maxSubarraySum(arr, 0)
	fmt.Println("Max subarray (any size):", maxVar)

	// Fixed window
	maxFixed, _ := maxSubarraySum(arr, 3)
	fmt.Println("Max subarray (size 3):", maxFixed)
}
