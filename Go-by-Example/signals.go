// sometimes we'd like our go programs to intelligently handle unix signals
// for example, we might want a server to gracefully shutdown when it receives SIGTERM, or a command-line tool to stop processing input if it receives SIGINT
// here's a modern way to handle signals using context
package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
)

func main() {
	// signal.NotifyContext returns a context that's canceled then one of the listed signals arrives
	ctx, stop := signal.NotifyContext(
		context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// this program will wait here until one of the configured signals is received
	fmt.Println("awaiting signal")
	<-ctx.Done()

	// context.Cause reports why the context was canceled
	// for a signal-triggered cancellation, this includes the signal value
	fmt.Println()
	fmt.Println(context.Cause(ctx))
	fmt.Println("exiting")
}

// when we run this program it will block waiting for a signal
// by typing ctrl-C (which the terminal shows as ^C) we can send a SIGINT signal, causing the program to print the cancellation and then exit
