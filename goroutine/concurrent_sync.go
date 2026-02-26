package main

import (
	"fmt"
	"sort"
	"sync"
)

func square(wg *sync.WaitGroup, ch <-chan int, result chan<- int) {
	defer wg.Done()
	for n := range ch {
		square := n * n
		result <- square
	}
}

func ConcurrentSync() {
	var num20 []int
	for i := 1; i <= 20; i++ {
		num20 = append(num20, i)
	}
	//fmt.Println(num20)

	wg := sync.WaitGroup{}

	ch := make(chan int, len(num20))
	result := make(chan int, len(num20))

	for range 5 {
		wg.Add(1)
		go square(&wg, ch, result)
	}

	for _, n := range num20 {
		ch <- n
	}
	close(ch)

	wg.Wait()
	close(result)

	var resNum20 []int
	for m := range result {
		resNum20 = append(resNum20, m)
	}

	sort.Ints(resNum20)
	fmt.Println("Sorted 20 Number :", resNum20)
}
