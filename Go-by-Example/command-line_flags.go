// command-line flags are a common way to specify options for command-line programs
// for example, in wc -l the -l is a command-line flag
package main

// go provides a flag package supporting basic command-line flag parsing
// we'll use this package to implement our example command-line program
import (
	"flag"
	"fmt"
)

func main() {
	// basic flag declarations are available for string, integer, and boolean options
	// here we declare a string flag word with a default value "foo" and a short description
	// this flag.String function returns a string pointer (not a string value); we'll see how to use this pointer below
	wordPtr := flag.String("word", "foo", "a string")

	// this declares numb and fork flags, using a similar approach to the word flag
	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	// it's also possible to declare an option that uses an existing var declared elsewhere in the program
	// not that we need to pass in a pointer to the flag declaration function
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// once all flags are declared, call flag.Parse() to execute the command-line parsing
	flag.Parse()

	// here we'll just dump out the parsed options and any trailing positional arguments
	// note that we need to dereference the pointers with e.g. *wordPtr to get the actual option values
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

// to experiment with the command-line flags program it's best to first compile it and then run the resulting binary directly
// if we omit flags they automatically take their default values
// trailling positional arguments can be provided after any flags
// the flag package requires all flags to appear before positional arguments (otherwise the flags will be interpreted as positional arguments)
// use -h or --help to get automatically generated help text for the command-line program
// if we provide a flag that wasn't specified to the flag package, the program will print an error message and show the help text again
