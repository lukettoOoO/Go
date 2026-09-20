// a panic typically means something went unexpectedly wrong
// mostly we use it to fail fast on errors that shouldn't occur during normal operation, or that we aren't prepared to handle gracefully
package main

import (
	"os"
	"path/filepath"
)

func main() {

	panic("a problem")

	// a common use of panic is to abort if a function returns an error value that we don't know how to (or want to) handle
	// here's an example of panicking if we get an unexpected error when creating a new file
	path := filepath.Join(os.TempDir(), "file")
	_, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	// running this program will cause it to panic, print an error message and goroutine traces, and exit with a non-zero status
	// when first panic in main fires, the program exits without reaching the rest of the code
	// note that unlike some languages which ese exceptions fot handling of many errors, in go it is idiomatic to use error-indicating return values whenever possible
}
