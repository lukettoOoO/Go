// Go supports constants of character, string, boolean, and numeric values
package main

import (
	"fmt"
	"math"
)

const s string = "constant" // const declares a constant value

// iota:
type Weekday int
const (
	Sunday Weekday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(s)

	const n = 500000000 // a const statement can also appear inside a function body

	const d = 3e20 / n // constant expressions perform arithmetic with arbitrary precision
	fmt.Println(d)

	fmt.Println(int64(d)) // a numeric constant has no type until it's given one, such as by an explicit conversion

	fmt.Println(math.Sin(n)) // a number can be given a type by using it in a context that requires one, such as a variable assignment or function call; for example, here math.Sin expects a float64

	fmt.Println(Sunday)
	fmt.Println(Friday)
}
