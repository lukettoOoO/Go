// some command-line tools, like the go tool or git have many subcommands, each with its own set of flags
// for example, go build and go get are two different subcommands of the go tool
// the flag package lets us easily define simple subcommands that have their own flags
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// we declare a subcommand using the NewFlagSet function, and proceed to define new flags specific for this subcommand
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	// for a different subcommand we can define different supported flags
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// the subcommand is expected as the first argument to the program
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// check which subcommand is invoked
	switch os.Args[1] {
		// for every subcommand, we parse its own flags and have access to trailling positional arguments
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
