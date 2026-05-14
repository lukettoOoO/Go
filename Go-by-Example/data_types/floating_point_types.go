package main

import (
	"fmt"
)

func main() {
	a, b := 5.67, 8.97
	fmt.Printf("type of a %T b %T\n", a, b)
	sum := a + b
	diff := a - b
	fmt.Printf("sum of %f and %f is %f, diff is %f\n", a, b, sum, diff)

	no1, no2 := 56, 89
	fmt.Printf("type of no1 %T no2 %T\n", no1, no2)
	fmt.Printf("sum of %d and %d is %d, diff is %d", no1, no2, no1 + no2, no1 - no2)
}
