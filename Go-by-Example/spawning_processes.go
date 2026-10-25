// sometimes our go programs need to spawn other processes
package main

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
)

func main() {
	// we'll start with a simple command that takes no arguments or input and just prints something to stdout
	// the exec.Command helper creates an object to represent this external process
	dateCmd := exec.Command("date")

	// the Output method runs the command, waits for it to finish and collects its standard output
	// if there were no errors, dateOut will hold bytes with the date info
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// Output and other methods of Command will return *exec.Error if there was a problem executing the command (e.g. wrong path), and *exec.ExitError if the command ran but exited with a non-zero return code
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		if e, ok := errors.AsType[*exec.Error](err); ok {
			fmt.Println("failed executing:", e)
		} else if e, ok := errors.AsType[*exec.ExitError](err); ok {
			exitCode := e.ExitCode()
			fmt.Println("command exit rc =", exitCode)
		} else {
			panic(err)
		}
	}

	// next we'll look at a slightly more involved case where we pipe data to the external process on its stdin and collect the results from its stdout
	grepCmd := exec.Command("grep", "hello")

	// here we explicitly grab input/output pipes, start the process, write some input to it, read the resulting output, and finally wait for the process to exit
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello hrep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	// we omitted error checks in the above example, but we could use the usual if err != nil pattern for all of them
	// we also only collect the StdoutPipe results, but we could collect the StderrPipe in exactly the same way
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// note that when spawning commands we need to provide an explicitly delineated command and argument array, vs. being able to just pass in one command-line string
	// if we want to spawn a full command with a string, we can use bash's -c option
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}

// the spawned programs return output that is the same as if we had run them directly from the command-line
// date doesn't habe a -x flag so it will exit with an error message and non-zero return code

