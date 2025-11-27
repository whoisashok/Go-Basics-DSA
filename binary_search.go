package main

import "fmt"

func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2
		midValue := arr[mid]
		if midValue == target {
			return mid
		} else if midValue < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func binarySearchManual() {
	sortedsearchList := []int{1, 3, 5, 7, 9, 11, 13, 14, 15}
	target := 7

	if index := BinarySearch(sortedsearchList, target); index != -1 {
		fmt.Printf("Element %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Element %d not found in the list\n", target)
	}
}
