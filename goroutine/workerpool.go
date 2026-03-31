package main

import "fmt"

const (
	numWorkers = 5
	numJobs    = 10
)

// Worker function that processes jobs and sends results
func poolWorker(id int, jobs <-chan int, results chan<- string) {
	for job := range jobs {
		results <- fmt.Sprintf("worker %d processed job %d", id, job)
	}
}

func WorkerPool() {
	// Create channels for jobs and results, with buffer to prevent blocking
	jobs := make(chan int, numWorkers)
	results := make(chan string, numWorkers)

	// Start worker pool
	for w := range numWorkers {
		go poolWorker(w, jobs, results)
	}

	// Send jobs to the workers
	for i := range numJobs {
		jobs <- i
	}
	// Close the jobs channel to indicate no more jobs will be sent
	close(jobs)

	// Collect all the results from the results channel
	for range numJobs {
		fmt.Println(<-results)
	}
	// Close the results channel after all results have been collected
	close(results)
}
