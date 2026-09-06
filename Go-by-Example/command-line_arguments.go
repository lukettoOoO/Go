// command-line arguments are a common way to parameterize execution of programs
// for example, go run hello.go uses run and hello.go arguments to the go program
package main

import (
	"fmt"
	"os"
)

func main() {

	// os.Args provides access to raw command-line arguments
	// note that the first value in this slice is the path to the program, and os.Args[1:] holds the arguments to the program
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// you can get individual args with normal indexing
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

	// to exepriment with command-line arguments it's best to build a binart with go build first
}
