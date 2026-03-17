package main

type Node struct {
	Data int
	Next *Node
}

type linkedList struct {
	Head *Node
}

func NewNode(data int) *Node {
	return &Node{Data: data, Next: nil}
}

func (list *linkedList) Append(data int) {
	node := NewNode(data)

	if list.Head == nil {
		list.Head = node
	} else {
		last := list.Head
		for last.Next != nil {
			last = last.Next // find the node at the end
		}
		last.Next = node // set the new node as the last
	}
}

func (list *linkedList) AddFirst(data int) {
	node := NewNode(data)

	if list.Head != nil {
		node.Next = list.Head // set head to the added value
	}
	list.Head = node
}

func (list *linkedList) Delete(data int) {
	if list.Head == nil {
		return
	}

	if list.Head.Data == data {
		list.Head = list.Head.Next
		return
	}

	current := list.Head
	for current.Next != nil {
		if current.Next.Data == data {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

func (list *linkedList) Display() []int {
	var elements []int
	current := list.Head

	for current != nil {
		elements = append(elements, current.Data)
		current = current.Next
	}

	return elements
}
func LinkedListDemo() {
	list := linkedList{}
	list.Append(20)
	list.AddFirst(5)
	list.Delete(20)
	elements := list.Display()
	for _, v := range elements {
		println(v)
	}
}
