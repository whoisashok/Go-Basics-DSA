package main

import (
	"fmt"
	"time"
)

func odd() {
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println("Odd	=	", i)
		}
	}
}

func even() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Even	=	", i)
		}
	}
}

func OddEvenGoroutine() {
	go odd()
	go even()
	time.Sleep(2 * time.Second)
}
