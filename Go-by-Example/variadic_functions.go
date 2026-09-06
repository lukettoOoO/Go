// variadic functions can be called with any number of trailing arguments
// for example, fmt.Println is a common variadic function

package main

import "fmt"

// here's a function that will take an arbitrary number of ints as arguments
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	// within a function, the type of nums is equivalent to []int
	// we can call len(nums), iterate over it with range etc.
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	// variadic functions can be called in the usual way with individual arguments
	sum(1, 2)
	sum(1, 2, 3)

	// if you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
