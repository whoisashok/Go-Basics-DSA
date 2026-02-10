package main

import "fmt"

func Segregate() {
	arr := []int{0, 1, 0, 1, 0, 0, 1, 1, 1, 0}

	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			arr[i], arr[j] = arr[j], arr[i]
			j++
		}
	}

	fmt.Println(arr)
}
