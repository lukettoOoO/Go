// go makes it possible to recover from a panic, by using the recover built-in function
// a recover can stop a panic from aborting the program and let it continue with execution instead

// an example of where this can be useful: a server wouldn't want to crash if one of the client connections exhibits a critical error
// instead, the server would want to close that connection and continue serving other clients
// in fact, this is what go's net/http does by default for http servers
package main

import "fmt"

// this function panics
func mayPanic() {
	panic("a problem")
}

// recover must be called within a deferred function
// when the enclosing function panics, the defer will activate and a recover call within it will catch the panic
func main() {

	// the return value of recover is the error raised in the call to panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	// this code will not run, because mayPanic panics
	// the execution of main stops at the point of the panic and resumes in the deferred closure
}
