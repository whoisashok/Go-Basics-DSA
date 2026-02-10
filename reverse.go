package main

import "fmt"

func reverseSlice[T any](s []T) {
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		s[left], s[right] = s[right], s[left]
	}
}

func reverseString(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func reverse() {
	intSlice := []int{1, 2, 3, 4, 5}
	reverseSlice(intSlice)
	fmt.Println("Reversed int slice:", intSlice) // Output: Reversed int slice: [5 4 3 2 1]

	stringSlice := []string{"apple", "banana", "cherry"}
	reverseSlice(stringSlice)
	fmt.Println("Reversed string slice:", stringSlice) // Output: Reversed string slice: [cherry banana apple]

	originalString := "Hello, World!"
	reversedString := reverseString(originalString)
	fmt.Println("Reversed string:", reversedString) // Output: Reversed string: !dlroW ,olleH

}
