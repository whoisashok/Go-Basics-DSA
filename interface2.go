package main

import "fmt"

// ArrayOps defines add and concat operations for a slice.
type ArrayOps[T any] interface {
	Add(elem T)
	Concat(other []T)
	Slice() []T
}

// GenericArray is a simple implementation of ArrayOps.
type GenericArray[T any] struct {
	elems []T
}

func NewGenericArray[T any](initial ...T) *GenericArray[T] {
	a := &GenericArray[T]{}
	if len(initial) > 0 {
		a.elems = append([]T{}, initial...)
	}
	return a
}

func (a *GenericArray[T]) Add(elem T) {
	a.elems = append(a.elems, elem)
}

func (a *GenericArray[T]) Concat(other []T) {
	a.elems = append(a.elems, other...)
}

func (a *GenericArray[T]) Slice() []T {
	return a.elems
}

func Interfaces2() {
	// int example
	var ints ArrayOps[int] = NewGenericArray[int](1, 2, 3)
	ints.Add(4)
	ints.Concat([]int{5, 6})
	fmt.Println(ints.Slice()) // Output: [1 2 3 4 5 6]

	// string example
	strs := NewGenericArray[string]("a", "b")
	var sops ArrayOps[string] = strs
	sops.Add("c")
	sops.Concat([]string{"d", "e"})
	fmt.Println(sops.Slice()) // Output: [a b c d e]
}
