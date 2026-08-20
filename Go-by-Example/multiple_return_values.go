// go has built-in support for multiple return values
// this feature is used in idiomatic go, for example to return both result and error values from a function
package main

import "fmt"

// the (int, int) in this function signature shows that the function returns 2 ints
func vals() (int, int) {
	return 6, 7
}

func main() {
	// here we use thw 2 different return values from the call with multiple assignment
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// if you only want a subset of the returned values, use the blank identifer _
	_, c := vals()
	fmt.Println(c)
}
