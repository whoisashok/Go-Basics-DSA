package main

// Generic type constraint to allow multiple types
type Element interface {
	int | string | any
}

// Generic function to print elements of a slice of any type
func printSlice[T Element](inputArray []T) {
	for _, v := range inputArray {
		println(v)
	}
}

// Main function to demonstrate generics usage
func Generic() {
	intSlice := []int{1, 2, 3, 4, 5}
	stringSlice := []string{"a", "b", "c", "d", "e"}

	printSlice(intSlice)
	printSlice(stringSlice)
}
