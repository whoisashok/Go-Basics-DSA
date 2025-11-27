package main

import "fmt"

type ArrayManipulation interface {
	ArrayAdd() []int
	ArrayConcat() []int
}

// Since the Arrayes type implements both ArrayAdd() and ArrayConcat() with the correct signatures,
// it implicitly implements the ArrayManipulation interface
type Arrayes struct {
	Arr1 []int
	Arr2 []int
}

func (a Arrayes) ArrayAdd() (res []int) {
	for i := range a.Arr1 {
		sum := a.Arr1[i] + a.Arr2[i]
		res = append(res, sum)
	}
	return res
}

func (a Arrayes) ArrayConcat() (res []int) {
	res = append(a.Arr1, a.Arr2...)
	return res
}

func Interfaces() {
	var arr1 = []int{1, 2, 3}
	var arr2 = []int{1, 2, 3}

	ArrayManipulationInterface := Arrayes{Arr1: arr1, Arr2: arr2}

	arrayAdd := ArrayManipulationInterface.ArrayAdd()
	arrayConcat := ArrayManipulationInterface.ArrayConcat()

	fmt.Println("arrayAdd    =    ", arrayAdd)
	fmt.Println("arrayConcat    =    ", arrayConcat)
}
