// the go standard library provides straightforward tools for outputting logs from go programs, with the log package for free-form output and the log/slog package for structured output
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"log/slog"
)

func main() {
	// simply invoking functions like Println from the log package uses the standard logger, which is already pre-configured for reasonable logging output to os.Stderr
	// additional methods like Fatal* or Panic* will exit the program after logging
	log.Println("standard logger")

	// loggers can be configured with flags to set their output format
	// by default, the standard logger has the log.Ldate and log.Ltime flags set, and these are collected in log.LstdFlags
	// we can chagne its flags to emit time with microsecond accuracy, for example
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	// it also supports emitting the file name and line from which the log function is called
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	// it may be useful to create a custom logger and pass it around
	// when creating a new logger, we can set a prefix to distinguish its output from other loggers
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")
}
