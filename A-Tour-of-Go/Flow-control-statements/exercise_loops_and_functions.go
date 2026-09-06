package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	prev_z := 0.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2*z)
		fmt.Printf("Iteration %d: %g\n", i, z)
		if i > 1 && math.Abs(prev_z - z) < 1.0 {
			break;
		}
		prev_z = z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(1024))
}
