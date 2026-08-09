package main

import (
	"golang.org/x/tour/tree"
	"fmt"
	"slices"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	s1 := make([]int, 0, 10)
	s2 := make([]int, 0, 10)
	
	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()

	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()
	
	for i := range ch1 {
		//fmt.Println(i)
		s1 = append(s1, i)
	}
	slices.Sort(s1)

	for i := range ch2 {
		//fmt.Println(i)
		s2 = append(s2, i)
	}
	slices.Sort(s2)

	if len(s1) != len(s2) {
		return false
	} else {
		for i := range len(s1) {
			if(s1[i] != s2[i]) {
				return false
			}
		}	
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
