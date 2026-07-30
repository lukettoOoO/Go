package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type
type List[T any] struct {
	next *List[T]
	val T
}

func (list *List[T]) ShowList() {
	for node := list; node != nil; node = node.next {
		fmt.Println(node.val)
	}
}

func main() {
	head := &List[int]{nil, 3}
	node := &List[int]{nil, 4}
	tail := &List[int]{nil, 5}

	head.next = node
	node.next = tail

	head.ShowList()
}
