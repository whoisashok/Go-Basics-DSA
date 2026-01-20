package main

import "fmt"

func longestSubstringNoRepeat(s string) (string, int) {
	lastSeen := make(map[byte]int)

	// Window is empty at the start
	left := 0
	maxLen := 0
	startIndex := 0

	for right := 0; right < len(s); right++ {

		// If duplicate found inside the window
		if duplicateIndx, found := lastSeen[s[right]]; found && duplicateIndx >= left {
			// Move the start to the right of the last occurrence
			// We must remove the previous 'c' from the window.
			left = duplicateIndx + 1
		}

		// Update last seen position
		lastSeen[s[right]] = right

		// Update max length
		currentLen := right - left + 1
		if currentLen > maxLen {
			maxLen = currentLen
			startIndex = left
		}
	}

	return s[startIndex : startIndex+maxLen], maxLen
}

func LongestSubstringNoRepeat() {
	s := "conceptoftheday"
	sub, length := longestSubstringNoRepeat(s)

	fmt.Println("The longest string is:", sub)              // Output: oftheday
	fmt.Println("The length of longest string is:", length) // Output: 8
}
