package main

import "fmt"

// removeDuplicates takes a slice of integers and returns a new slice
// containing only the unique elements.
func removeIntDuplicates(input []int) []int {
	// Create a map to store encountered elements.
	// The keys are the elements, and the values are booleans (true if encountered).
	encountered := make(map[int]bool)
	var result []int // This will store the unique elements

	for _, v := range input {
		// If the element has not been encountered before, add it to the result slice
		// and mark it as encountered in the map.
		if !encountered[v] {
			encountered[v] = true
			result = append(result, v)
		}
	}
	return result
}

// removeStringDuplicates is a similar function for string slices
func removeStringDuplicates(input []string) []string {
	encountered := make(map[string]bool)
	var result []string
	for _, v := range input {
		if !encountered[v] {
			encountered[v] = true
			result = append(result, v)
		}
	}
	return result
}

// removeStringDuplicates is a similar function for strings to remove duplicate characters
func removeDuplicateChars(input string) []string {
	encountered := make(map[rune]bool)
	var result []string
	for _, v := range input {
		if !encountered[v] {
			encountered[v] = true
			result = append(result, string(v))
		}
	}
	return result
}

func RemoveDuplicates() {
	numbers := []int{1, 2, 2, 3, 4, 4, 5, 1, 6}
	fmt.Println("Original slice:", numbers)

	uniqueNumbers := removeIntDuplicates(numbers)
	fmt.Println("Slice after removing duplicates:", uniqueNumbers)

	// Example with strings
	words := []string{"apple", "banana", "apple", "orange", "banana"}
	fmt.Println("Original string slice:", words)
	uniqueWords := removeStringDuplicates(words) // Assuming a similar function for strings
	fmt.Println("String slice after removing duplicates:", uniqueWords)

	inputString := "hello world"
	uniqueString := removeDuplicateChars(inputString)
	fmt.Printf("Original string: %s\n", inputString)
	fmt.Printf("String with unique characters: %s\n", uniqueString) // Output: helo wrd
}
