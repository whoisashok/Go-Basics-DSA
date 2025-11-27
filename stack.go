package main

import (
	"fmt"
)

// Stack represents a stack using a slice
type Stack struct {
	elements []string
}

// Push adds an element to the top of the stack
func (s *Stack) Push(element string) {
	s.elements = append(s.elements, element) // Appends the element to the slice
}

// Pop removes the top element from the stack
func (s *Stack) Pop() string {
	if len(s.elements) == 0 {
		return "No elements to remove!"
	}
	topElement := s.elements[len(s.elements)-1] // Get the last element
	s.elements = s.elements[:len(s.elements)-1] // Remove the last element
	return topElement
}

// Top returns the top element of the stack without removing it
func (s *Stack) Top() string {
	if len(s.elements) == 0 {
		return "No elements in the stack!"
	}
	return s.elements[len(s.elements)-1]
}

func StackList() {
	stack := Stack{}
	stack.Push("Element 1") // Adding first element
	stack.Push("Element 2") // Adding second element

	// Checking the top element, which is "Element 2"
	fmt.Println("Top element:", stack.Top()) // Outputs: Top element: Element 2

	// Removing the top element, which is "Element 2"
	fmt.Println("Removed:", stack.Pop()) // Outputs: Removed: Element 2
}
