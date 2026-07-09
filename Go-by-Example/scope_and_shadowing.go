package main

import "fmt"

func main() {
	
	x := 10 // outer variable
	fmt.Println(x) // prints: 10

	if true {
		x := 5 // inner variable (shadows the outer x)
		fmt.Println(x) // prints: 5 (inner x)
	}

	fmt.Println(x) // prints: 10 (outer x, accessible again)
}
