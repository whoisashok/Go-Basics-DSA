package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/sony/gobreaker"
)

func main() {
	settings := gobreaker.Settings{
		Name:        "API-CB",
		MaxRequests: 3,
		Interval:    time.Second * 10,
		Timeout:     time.Second * 5,
	}

	cb := gobreaker.NewCircuitBreaker(settings)

	for i := 0; i < 10; i++ {
		result, err := cb.Execute(func() (interface{}, error) {
			// simulate failure
			return nil, errors.New("service down")
		})

		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Println(result)
	}
}
