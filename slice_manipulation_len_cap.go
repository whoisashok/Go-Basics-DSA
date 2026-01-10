package main

import "fmt"

func SliceManipulation() {

	ar := []int{1, 2, 3, 4, 5}        // len= 5, cap = 5
	fmt.Println(ar, len(ar), cap(ar)) // 1, 2, 3, 4, 5		len 5 cap 5

	ar = append(ar, 100) // 1, 2, 3, 4, 5 ,100		len 6 cap = 10
	fmt.Println(ar, len(ar), cap(ar))

	ar1 := ar[:] // 1, 2, 3, 4, 5 ,100 // len 6 cap = 10
	fmt.Println(ar1, len(ar1), cap(ar1))

	ar1 = append(ar1, 200)               // 1, 2, 3, 4, 5 ,100, 200 // len 7 cap = 10
	fmt.Println(ar1, len(ar1), cap(ar1)) // 1, 2, 3, 4, 5 ,100, 200 // len 7 cap = 10

	ar1[0] = 200                         // 200, 2, 3, 4, 5 ,100, 200
	fmt.Println(ar1, len(ar1), cap(ar1)) // 200, 2, 3, 4, 5 ,100, 200 // len 7 cap = 10
}
