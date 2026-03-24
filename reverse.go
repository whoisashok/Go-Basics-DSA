package main

import (
	"fmt"
	"strings"
)

// reverseWords reverses the order of words in a sentence.
func reverseWords(s string) string {
	// Split the string into a slice of words using space as a delimiter
	words := strings.Fields(s) // strings.Fields handles multiple spaces and trims leading/trailing spaces

	// Reverse the order of elements in the slice using a two-pointer approach
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Join the reversed words back into a single string with spaces
	return strings.Join(words, " ")
}

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

	sentence := "the sky is blue"
	reversedSentence := reverseWords(sentence)
	fmt.Printf("Original sentence: \"%s\"\n", sentence)
	fmt.Printf("Reversed sentence: \"%s\"\n", reversedSentence)

}
