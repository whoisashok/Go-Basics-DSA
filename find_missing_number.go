package main

import "fmt"

func findMissing(nums []int) int {
	n := len(nums) + 1
	seen := make(map[int]bool)

	for _, num := range nums {
		seen[num] = true
	}

	for i := 1; i <= n; i++ {
		if !seen[i] {
			return i
		}
	}
	return -1
}

func FindMissing() {
	nums := []int{1, 3, 0, 4}
	fmt.Println(findMissing(nums))
}
