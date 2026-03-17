package main

import (
	"fmt"
	"sync"
)

type Job struct {
	Index int
	Value int
}

type Result struct {
	Index int
	Value int
}

func square(wg *sync.WaitGroup, jobs <-chan Job, results chan<- Result) {
	defer wg.Done()
	for job := range jobs {
		results <- Result{
			Index: job.Index,
			Value: job.Value * job.Value,
		}
	}
}

func main() {
	const total = 20

	jobs := make(chan Job, total)
	results := make(chan Result, total)

	var wg sync.WaitGroup

	// Start 5 workers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go square(&wg, jobs, results)
	}

	// Send jobs with index
	for i := 1; i <= total; i++ {
		jobs <- Job{
			Index: i - 1,
			Value: i,
		}
	}
	close(jobs)

	wg.Wait()
	close(results)

	// Pre-allocate slice
	res := make([]int, total)

	// Place result in correct index
	for r := range results {
		res[r.Index] = r.Value
	}

	fmt.Println("Squares in order:", res)
}
