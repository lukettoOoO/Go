package main

// complex64 - complex numbers with float32 real and imaginary parts
// complex128 - complex numbers with float64 real and imaginary parts

import (
	"fmt"
)

func main() {
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum: ", cadd)
	cmul := c1 * c2
	fmt.Println("product: ", cmul)
}
