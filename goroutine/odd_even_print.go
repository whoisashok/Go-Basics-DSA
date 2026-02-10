package main

import (
	"fmt"
	"sync"
)

const limit = 100

func oddFunc(oddChan, evenChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= limit; i += 2 {
		<-oddChan
		fmt.Println(i)
		evenChan <- struct{}{}
	}
}

func evenFunc(oddChan, evenChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= limit; i += 2 {
		<-evenChan
		fmt.Println(i)
		if i < limit {
			oddChan <- struct{}{}
		}
	}
}

func OddEvenGoroutine() {

	oddChan := make(chan struct{})
	evenChan := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(2)

	go oddFunc(oddChan, evenChan, &wg)
	go evenFunc(oddChan, evenChan, &wg)

	// Start with odd number
	oddChan <- struct{}{}

	wg.Wait()
	fmt.Println("Done!")
}
