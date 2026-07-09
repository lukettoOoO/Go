package main

import "fmt"

func main() {
	a := true
	var b bool = false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)
}
