package main

import "fmt"

func TopKFrequent() {
	r := topKFrequent([]int{3, 3, 3, 1, 1, 2, 0, 0, 0, 0, 0}, 2)
	fmt.Println(r)
}

func topKFrequent(nums []int, k int) []int {
	freq := map[int]int{}
	for _, v := range nums {
		freq[v]++
	}
	fmt.Println(freq) // map[1:3 2:2 3:1]

	//Create a 2D slice [][]int of size len(nums)+1, where the index represents the frequency.
	buckets := make([][]int, len(nums)+1)
	for num, f := range freq {
		// Append the number to the corresponding frequency bucket
		buckets[f] = append(buckets[f], num)
	}
	fmt.Println(buckets) // [[] [3] [2] [1] [] [] []]

	res := []int{}
	// Iterate through the buckets backwards (from highest frequency to lowest) until you have elements.
	for i := len(buckets) - 1; i >= 0 && len(res) < k; i-- {
		res = append(res, buckets[i]...)
	}
	return res[:k]
}
