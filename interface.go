package main

import "fmt"

// Interface definition
type ArrayOperations interface {
	ArrayAdd() []int
	ArrayConcat() []int
}

// Since the ArrayImpl type implements both ArrayAdd() and ArrayConcat() with the correct signatures,
// it implicitly implements the ArrayOperations interface
type ArrayImpl struct {
	Arr1 []int
	Arr2 []int
}

// Method for element-wise array addition
func (a ArrayImpl) ArrayAdd() (result []int) {
	if len(a.Arr1) != len(a.Arr2) {
		fmt.Println("Arrays must be of same length for addition")
		return nil
	}
	for i := range a.Arr1 {
		sum := a.Arr1[i] + a.Arr2[i]
		result = append(result, sum)
	}
	return result
}

// Method for array concatenation
func (a ArrayImpl) ArrayConcat() (result []int) {
	result = append(a.Arr1, a.Arr2...)
	return result
}

// Main function to demonstrate interface usage
func Interfaces() {
	var arr1 = []int{1, 2, 3}
	var arr2 = []int{1, 2, 3}

	ArrayOperationsInterface := ArrayImpl{Arr1: arr1, Arr2: arr2}

	arrayAdd := ArrayOperationsInterface.ArrayAdd()
	arrayConcat := ArrayOperationsInterface.ArrayConcat()

	fmt.Println("Array 1:", arr1)
	fmt.Println("Array 2:", arr2)
	fmt.Println("Array Add:", arrayAdd)
	fmt.Println("Array Concat:", arrayConcat)
}
