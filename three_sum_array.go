package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int, targetSum int) [][]int {
	sort.Ints(nums)
	pairs := [][]int{}

	// because we are picking 3 numbers: i, left, right
	// so we need to go until len(nums) - 2
	// -2 ensures there are always two numbers left to form a triplet.
	// panic: index out of range
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1

		for left < right {
			currentSum := nums[i] + nums[left] + nums[right]

			if currentSum == targetSum {
				pairs = append(pairs, []int{nums[i], nums[left], nums[right]})
				left++
				right--
			} else if currentSum < targetSum {
				left++
			} else {
				right--
			}
		}
	}

	return pairs
}

func threeSumArray() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	targetSum := 0
	fmt.Println(threeSum(nums, targetSum))
}
