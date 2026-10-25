// use os.Exit to immediately exit with a given status
package main

import (
	"fmt"
	"os"
)

func main() {
	// defers will not be run when using os.Exit, so this fmt.Println will never be called
	defer fmt.Println("!")
	// exit with status 3
	os.Exit(3)
}

// note that unlike e.g. C, go does not use an integer return value from main to indicate exit status
// if we'd like to exit with a non-zero status we should use os.Exit
// if we run exit.go using go run, the exit will be picked up by go and printed
// by building and executing a binary we can see the status in the terminal
// note that the ! from our program never got printed
