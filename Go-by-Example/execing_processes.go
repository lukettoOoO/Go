// sometimes we just want to completely replace the current go process with another (perhaps a non-go) one
// to do this we'll use go's implementation of the classic exec function
package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// for our example we'll exec ls
	// go requires an absolute path to the binary we want to execute, so we'll use exec.LookPath to find it (probably /bin/ls)
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// Exec requires arguments in slice form (as opposed to one big string)
	// we'll give ls a few common arguments
	// note that the first argument should be the program name
	args := []string{"ls", "-a", "-l", "-h"}

	// Exec also needs a set of environment variables to use
	// here we just provide our current environment
	env := os.Environ()

	// here's the actual syscall.Exec call
	// if this call is successful, the execution of our process will end here and be replaced by the /bin/ls -a -l -h process
	// if there is an error we'll get a return value
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

// note that go does not offer a classic Unix fork function
// usually this isn't an issue though, since starting goroutines, spawning processes, and exec'ing processes covers most use cases for fork
