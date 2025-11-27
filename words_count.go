package main

import (
	"fmt"
	"strings"
)

func WordsCount() {
	str := "betty bought the butter , the butter was bitter , " +
		"betty bought more butter to make the bitter butter better "
	words := strings.Split(str, " ")
	wordCount := make(map[string]int)

	for _, word := range words {
		_, exists := wordCount[word]
		if exists {
			wordCount[word] += 1
		} else {
			wordCount[word] = 1
		}
	}

	for key, value := range wordCount {
		fmt.Printf("%s: %d\n", key, value)
	}
}
