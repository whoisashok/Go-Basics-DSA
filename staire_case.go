package main

import (
	"fmt"
	"strings"
)

func staircase(n int) {
	for i := 1; i <= n; i++ {
		// Print spaces for right alignment
		fmt.Print(strings.Repeat(" ", n-i))
		// Print hash symbols for the current step
		fmt.Print(strings.Repeat("#", i))
		// Move to the next line
		fmt.Println()
	}
}

func Staircase() {
	staircase(4)
}
