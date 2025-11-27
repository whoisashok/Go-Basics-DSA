package main

import (
	"fmt"
	"strings"
)

func long_sub_string_norepeat_chars() {
	// Define the string to be tested
	var str string = "conceptoftheday"

	// Define variables to store the longest known string
	// and the current test string
	var longestString string = ""
	var currentString string = ""

	// Loop through each of the characters in the string
	for i := 0; i < len(str); i++ {
		// Get the next character
		var char string = string(str[i])

		// Check to see if the current test string contains
		// the next character
		if strings.Contains(currentString, char) {

			// See if the string we have finished
			// finding is the longest string
			if currentString > longestString {
				longestString = currentString
			}

			// Split the string on the repeated character
			currentString = strings.Split(currentString, char)[1]
		}

		// Add the next character to our current string
		currentString = currentString + char
	}

	// Final check at the end of the loop
	if len(currentString) > len(longestString) {
		longestString = currentString
	}

	// Output the longest string
	fmt.Println("The longest string is:", longestString)
	fmt.Println("The length of longest string is:", len(longestString))
}
