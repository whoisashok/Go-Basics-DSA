package main

import (
	"sync"
	"time"
)

var val int

func writeVal(mu *sync.RWMutex) {
	mu.Lock() // exclusive lock for writing
	defer mu.Unlock()

	val++
	println("write:", val)
}

func readVal(mu *sync.RWMutex) {
	mu.RLock() // shared lock for reading
	defer mu.RUnlock()

	println("read:", val)
}

func MutexLocks() {
	var mu sync.RWMutex

	// writers
	for i := 0; i < 5; i++ {
		go writeVal(&mu)
	}

	// readers
	for i := 0; i < 5; i++ {
		go readVal(&mu)
	}

	time.Sleep(time.Second)
}
