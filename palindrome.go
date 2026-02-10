package main

import (
	"fmt"
	"strings"
)

func isPalindromeString(s string) bool {
	// Convert to lowercase and remove spaces for case-insensitive and space-agnostic comparison
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))

	// Use two pointers, one from the start and one from the end
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		if s[left] != s[right] {
			return false
		}
	}
	return true
}

func isPalindromeNumber(n int) bool {
	if n < 0 { // Negative numbers are not palindromes
		return false
	}

	originalNum := n
	reversedNum := 0

	for n > 0 {
		remainder := n % 10
		reversedNum = reversedNum*10 + remainder
		n /= 10
	}

	return originalNum == reversedNum
}

func Palindrome() {
	fmt.Println(isPalindromeString("madam"))                       // true
	fmt.Println(isPalindromeString("A man a plan a canal Panama")) // true
	fmt.Println(isPalindromeString("hello"))                       // false

	fmt.Println(isPalindromeNumber(121))  // true
	fmt.Println(isPalindromeNumber(123))  // false
	fmt.Println(isPalindromeNumber(-121)) // false
}
