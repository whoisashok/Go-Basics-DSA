package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func workerWithContext(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context worker stopped:", ctx.Err())
			return
		default:
			fmt.Println("context worker running")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func workerWithDone(done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-done:
			fmt.Println("done-channel worker stopped")
			return
		default:
			fmt.Println("done-channel worker running")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func workerWithQuit(quit <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case msg := <-quit:
			fmt.Println("quit-channel worker stopped:", msg)
			return
		default:
			fmt.Println("quit-channel worker running")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func workerWithAtomic(stop *atomic.Bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for !stop.Load() {
		fmt.Println("atomic worker running")
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("atomic worker stopped")
}

func workerNatural(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("natural worker started")
	time.Sleep(2 * time.Second)
	fmt.Println("natural worker finished")
}

func StopGoroutine() {
	var wg sync.WaitGroup

	// 1. Context-based cancellation
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go workerWithContext(ctx, &wg)

	// 2. Done channel
	done := make(chan struct{})
	wg.Add(1)
	go workerWithDone(done, &wg)

	// 3. Quit channel with message
	quit := make(chan string)
	wg.Add(1)
	go workerWithQuit(quit, &wg)

	// 4. Atomic flag
	var stop atomic.Bool
	wg.Add(1)
	go workerWithAtomic(&stop, &wg)

	// 5. Natural completion
	wg.Add(1)
	go workerNatural(&wg)

	// Let workers run
	time.Sleep(2 * time.Second)

	// Stop everything
	fmt.Println("\n--- stopping workers ---")

	// 1. Context-based cancellation
	cancel()
	// 2. Done channel
	close(done)
	// 3. Quit channel with message
	quit <- "shutdown now"
	// 4. Atomic flag
	stop.Store(true)

	wg.Wait()
	fmt.Println("\nall goroutines stopped cleanly")
}
