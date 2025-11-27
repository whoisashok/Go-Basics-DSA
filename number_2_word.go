package main

import (
	"fmt"
	"strings"
)

var belowTwenty = []string{
	"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine",
	"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen",
}

var tens = []string{
	"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety",
}

var thousands = []string{
	"", "Thousand", "Million", "Billion", "Trillion",
}

// Function to convert a number less than 1000 to words
func convertThreeDigits(num int) string {
	if num == 0 {
		return ""
	}

	result := ""
	if num >= 100 {
		result += belowTwenty[num/100] + " Hundred "
		num %= 100
	}
	if num >= 20 {
		result += tens[num/10] + " "
		num %= 10
	}
	if num > 0 {
		result += belowTwenty[num] + " "
	}

	return strings.TrimSpace(result)
}

// Function to convert a number to words
func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	res := ""
	idx := 0
	for num > 0 {
		if num%1000 != 0 {
			res = convertThreeDigits(num%1000) + " " + thousands[idx] + " " + res
		}
		num /= 1000
		idx++
	}
	return strings.TrimSpace(res)
}

func NumberToWords() {
	num := 999999999
	result := numberToWords(num)
	fmt.Println("Number in words:", result)
}
