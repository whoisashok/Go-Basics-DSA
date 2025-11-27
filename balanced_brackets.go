package main

import (
	"fmt"
)

// isBalanced checks if a string has balanced brackets.
func isBalanced(s string) bool {
	// Use a slice as a stack to store opening brackets.
	var stack []rune

	// Define a map to store matching bracket pairs.
	bracketMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, char := range s {
		switch char {
		case '(', '[', '{': // If it's an opening bracket, push it onto the stack.
			stack = append(stack, char)
		case ')', ']', '}': // If it's a closing bracket.
			if len(stack) == 0 { // If the stack is empty, there's no matching opening bracket.
				return false
			}
			// Pop the last opening bracket from the stack.
			lastOpen := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Check if the popped opening bracket matches the current closing bracket.
			if bracketMap[lastOpen] != char {
				return false
			}
		}
	}

	// After iterating through the string, if the stack is empty, all brackets are balanced.
	return len(stack) == 0
}

func BalancedBrackets() {
	fmt.Println(isBalanced("{[()]}")) // true
	fmt.Println(isBalanced("{[(])}")) // false
	fmt.Println(isBalanced("((()))")) // true
	fmt.Println(isBalanced("(()"))    // false
	fmt.Println(isBalanced(")"))      // false
	fmt.Println(isBalanced(""))       // true
}
