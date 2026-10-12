// environment variables are a universal mechanism for conveying configuration infromation to Unix programs
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// to set a key/value pair, use os.Setenv
	// to get a value for a key, use os.Getenv
	// this will return an empty string if the key isn't present in the environment
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// use os.Environ to list all key/value pairs in the environment
	// this returns a slice of strings in the form KEY=value
	// we can strings.SplitN them to get the key and value
	// here we print all the keys
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}

// running the program shows that we pick up the value for FOO that we set in the program, but that BAR is empty
// the list of keys in the environment will depend on our particular machine
// if we set BAR in the environment first, the running program picks that value up
