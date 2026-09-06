package main

import (
	"fmt"
)

func main() {
	a := 80 // int
	b := 91.8 // float64
	sum := a + int(b) // int + float64 not allowed
	fmt.Println(sum)

	i := 10
	var j float64 = float64(i) // this statement will not work without explicit conversion
	fmt.Println("j =", j)
}
