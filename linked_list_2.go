package main

import (
	"container/list"
	"fmt"
)

func LinkedList() {
	fmt.Println("Go Linked Lists Tutorial")

	mylist := list.New()
	mylist.PushBack(1)
	mylist.PushFront(2)
	mylist.PushBack(3)
	mylist.InsertBefore(2, mylist.Front())
	mylist.InsertAfter(4, mylist.Back())

	// Remove an element
	element := mylist.Front() // Get the first element
	mylist.Remove(element)    // Remove the first element

	// Iterate through the list
	for element := mylist.Front(); element != nil; element = element.Next() {
		// do something with element.Value
		fmt.Println(element.Value)
	}

}
