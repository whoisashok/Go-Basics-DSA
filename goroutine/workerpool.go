package main

import (
	"fmt"
	"sync"
)

func workerSquare(wg *sync.WaitGroup, ch <-chan int, result chan<- int) {
	defer wg.Done()
	for n := range ch {
		square := n * n
		result <- square
	}
}

func WorkerPool() {
	var num20 []int
	// Generate numbers from 1 to 20
	for i := 1; i <= 20; i++ {
		num20 = append(num20, i)
	}
	//fmt.Println(num20)

	wg := sync.WaitGroup{}

	ch := make(chan int, len(num20))
	result := make(chan int, len(num20))

	// Start worker pool
	for range 5 {
		wg.Add(1)
		go workerSquare(&wg, ch, result)
	}

	// Send numbers to the channel
	for _, n := range num20 {
		ch <- n
	}
	close(ch)

	wg.Wait()
	close(result)

	// Collect results
	var resNum20 []int
	for m := range result {
		resNum20 = append(resNum20, m)
	}

	fmt.Println("Sorted 20 Number :", resNum20)
}
