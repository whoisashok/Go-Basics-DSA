package main

import (
	"fmt"
	"sync"
)

func sender(ch chan int, wg *sync.WaitGroup) {
	for i := range 10 {
		ch <- i
	}
	close(ch)
	wg.Done()
}

func receiver(ch chan int, wg *sync.WaitGroup) {
	for result := range ch {
		fmt.Println("Receiving from channel:", result)
	}
	wg.Done()
}

func SendReceiveChannels() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)
	go sender(ch, &wg)
	go receiver(ch, &wg)

	wg.Wait()
}
