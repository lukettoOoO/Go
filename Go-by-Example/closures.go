// go supports annonymous functions, which can form closures
// annonymous functions are useful when you want to define a function inline without having to name it

package main

import "fmt"

// this function intSeq returns another function, which we define annonymously in the body of intSeq
// the returned function closes over the variable i to form a closure

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()
	
	// we call intSeq, assigning the result (a function) to nextInt
	// this function value captures its own i value, which will be updated each time we call nextInt

	// see this effect of the closure by calling nextInt a few times
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// to confirm that the state is unique to that particular function, create and test a new one
	newInts := intSeq()
	fmt.Println(newInts())
}
