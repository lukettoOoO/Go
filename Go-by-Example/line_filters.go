// a line filter is a common type of program that reads input on stdin, processes it, and then prints some derived result to stdout
// grep and sed are common line filters
package main

// here's am example line filter in go that writes a capitalized version of all input text
// we can use this pattern to write our own go line filters
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// wrapping the unbuffered os.Stdin with a buffered scanner gives us a convenient Scan method that advances the scanner to the next token; which is the next line in the default scanner
	scanner := bufio.NewScanner(os.Stdin)

	// Text returns the current token, here the next line, from the input
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		// write out the uppercased line
		fmt.Println(ucl)
	}

	// check for errors during Scan
	// end of file is expected and not reported by Scan as an error
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	// to try out the line filter, first make a file with a few lowercase lines
	// then use the line filter to get uppercase lines
}
