package main

// rune is an alias of int32
// conceptually, it represents a unicode code point

import "fmt"

func main() {
	var c rune = 'a'
	r := 'G' // this is a rune literal (note the single quotes)
	fmt.Printf("Type: %T, Value: %v\n", r, r)
	fmt.Printf("%c\n", c)
}
