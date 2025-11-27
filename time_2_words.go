package main

import "fmt"

// timeInWords converts a given hour (h) and minute (m) into words.
// It handles special cases like quarter past, half past, and o'clock.
func timeInWords(h int32, m int32) string {
	// Array of words for numbers 0 through 29.
	// This covers all minute counts and all 12 hours.
	numWords := []string{
		"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
		"eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
		"twenty", "twenty one", "twenty two", "twenty three", "twenty four", "twenty five", "twenty six", "twenty seven", "twenty eight", "twenty nine",
	}

	// Calculate the word for the current hour (h)
	currentHourWord := numWords[h]

	// Determine the next hour (handles 12:XX wrapping to 1)
	nextHour := (h % 12) + 1
	nextHourWord := numWords[nextHour]

	switch m {
	case 0:
		// e.g., 5:00 -> five o' clock
		return fmt.Sprintf("%s o' clock", currentHourWord)
	case 15:
		// e.g., 5:15 -> quarter past five
		return fmt.Sprintf("quarter past %s", currentHourWord)
	case 30:
		// e.g., 5:30 -> half past %s
		return fmt.Sprintf("half past %s", currentHourWord)
	case 45:
		// e.g., 5:45 -> quarter to six (next hour)
		return fmt.Sprintf("quarter to %s", nextHourWord)
	default:
		if m < 30 {
			// Minutes past the hour (1-29, excluding 15)
			minuteWord := numWords[m]
			minuteUnit := "minutes"
			if m == 1 {
				minuteUnit = "minute"
			}
			// e.g., 5:02 -> two minutes past five
			return fmt.Sprintf("%s %s past %s", minuteWord, minuteUnit, currentHourWord)
		} else {
			// Minutes to the next hour (31-59, excluding 45)
			minutesTo := 60 - m
			minuteWord := numWords[minutesTo]
			minuteUnit := "minutes"
			if minutesTo == 1 {
				minuteUnit = "minute"
			}
			// e.g., 5:40 (20 minutes to 6) -> twenty minutes to six
			return fmt.Sprintf("%s %s to %s", minuteWord, minuteUnit, nextHourWord)
		}
	}
}

func TimeToWords() {
	// Example 1: 5:02
	fmt.Printf("5:02 -> %s\n", timeInWords(5, 2))

	// Example 2: 5:30 (half past)
	fmt.Printf("5:30 -> %s\n", timeInWords(5, 30))

	// Example 3: 7:45 (quarter to 8)
	fmt.Printf("7:45 -> %s\n", timeInWords(7, 45))

	// Example 4: 12:00 (o' clock)
	fmt.Printf("12:00 -> %s\n", timeInWords(12, 0))

	// Example 5: 12:58 (2 minutes to 1)
	fmt.Printf("12:58 -> %s\n", timeInWords(12, 58))
}
