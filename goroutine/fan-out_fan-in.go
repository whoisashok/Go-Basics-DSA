package main

import (
	"fmt"
	"sync"
)

const (
	noJobs    = 20
	noWorkers = 5
)

// ✅ Channels for jobs and results
var jobs = make(chan int, noJobs)
var results = make(chan int, noJobs)
var wg sync.WaitGroup
var resNum20 []int

// Worker (fan-out)
func worker() {
	defer wg.Done()
	for job := range jobs {
		results <- job
	}
}

// Worker pool (fan-out)
func FanOutFanInPool() {
	for range numWorkers {
		wg.Add(1)
		go worker()
	}
}

// Publisher (producer)
func publisher() {
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
}

// Subscriber (consumer)
func subscriber() {
	for result := range results {
		resNum20 = append(resNum20, result)
	}
}

func FanOutFanIn() {
	// ✅ Fan-out: start workers
	FanOutFanInPool()

	// ✅ Producer (can also be a goroutine for full async)
	go publisher()

	// ✅ Fan-in closer (important pattern)
	go func() {
		wg.Wait()
		close(results)
	}()

	// ✅ Fan-in consumer (async stream processing)
	subscriber()

	fmt.Println("Async fan-out fan-in result:", resNum20)
}
