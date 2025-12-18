package main

import "fmt"

func PrintNumbers(odd, even <-chan int) {
	for {
		select {
		case oddEvent := <-odd:
			fmt.Println(oddEvent, "odd channel")

		case evenEvent := <-even:
			fmt.Println(evenEvent, "even channel")
		}
	}
}

func OddEvenGoroutine() {
	// Creating channels to send odd and even numbers
	odd := make(chan int)
	even := make(chan int)
	// Start a goroutine to handle the printing
	go PrintNumbers(odd, even)
	// Send numbers to the channels
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			even <- i // Send even numbers to 'even' channel
		} else {
			odd <- i // Send odd numbers to 'odd' channel
		}
	}
	close(odd)
	close(even)
}
