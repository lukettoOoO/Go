// starting with version 1.23, Go has added support for iterators, which lets us range over pretty much anything
package main

import (
	"fmt"
	"iter"
	"slices"
	"strings"
)

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// all returns an iterator, which in go is a function with a special signature
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		// the iterator function takes another function as a parameter, called yield by convention (but the name can be arbitrary)
		// it will call yield for every element we want to iterate over, and note yield's return value for a potential early termination
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

// iteration doesn't require an underlying data structure, and doesn't even have to be finite
// here's a function returning an iterator over fibonacci numbers: it keeps running as long as yield keeps returning true
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	// since List.All returns an iterator, we can use it in a regular range loop
	for e := range lst.All() {
		fmt.Println(e)
	}

	// packages like slices have a number of useful functions to work with iterators
	// for example, Collect takes any iterator and collects all its values into a slice
	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	// standard library packages now expose iterator helpers too
	// for example, strings.SplitSeq iterates over parts of a byte slice without first building a result slice
	for part := range strings.SplitSeq("go-by-example", "-") {
		fmt.Printf("part: %s\n", part)
	}

	for n := range genFib() {
		// once the loop hits break or an early return, the yield function passed to the iterator will reuturn false
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
