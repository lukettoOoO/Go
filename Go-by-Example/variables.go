// in Go, variables are explicitly declared and used by the compiler to e.g. chack type-correctness of function calls
package main

import "fmt"

func main() {

	var a = "initial" // var declares 1 or more variables
	fmt.Println(a)

	var b, c int = 1, 2 // you can declare multiple variables at once
	fmt.Println(b, c)

	var d = true // Go will infer the type of initialized variables
	fmt.Println(d)

	var e int // variables declared without a corresponding intitialization are zero-valued; for example, the zero value for an int is 0
	fmt.Println(e)

	f := "apple"
	fmt.Println(f) // the := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case; this syntax is only available inside functions
}
