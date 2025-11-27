package main

import (
	"fmt"
)

// Queue represents a queue using a slice
type Queue struct {
	elements []string
}

// Enqueue adds an element to the end of the queue
func (q *Queue) Enqueue(element string) {
	q.elements = append(q.elements, element) // Appends the element to the slice
}

// Dequeue removes the first element from the queue
func (q *Queue) Dequeue() string {
	if len(q.elements) == 0 {
		return "No elements to dequeue!"
	}
	frontElement := q.elements[0] // Get the first element
	q.elements = q.elements[1:]   // Remove the first element
	return frontElement
}

// Front returns the first element of the queue without removing it
func (q *Queue) Front() string {
	if len(q.elements) == 0 {
		return "No elements in the queue!"
	}
	return q.elements[0]
}

func QueueList() {
	queue := Queue{}
	queue.Enqueue("Element 1") // Adds first element to the queue
	queue.Enqueue("Element 2") // Adds second element to the queue

	// Checking the front element, which is "Element 1"
	fmt.Println("Front element:", queue.Front()) // Outputs: Front element: Element 1

	// Removing the first element, which is "Element 1"
	fmt.Println("Removed:", queue.Dequeue()) // Outputs: Removed: Element 1
}
