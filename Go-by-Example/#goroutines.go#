// a goroutine is a lightweight thread of execution
package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// suppose we have a function call f(s)
	// here's how we'd call that in the usual way, running it synchronously
	f("direct")

	// to invoke this function in a goroutine, use go f(s)
	// this new goroutine will execute concurrently with the calling one
	go f("goroutine")

	// you can also start a goroutine for an annonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// out two function calls are running asynchronously in separate goroutines now
	// wait for them to finish (for a more robust apporach, use a WaitGroup)
	time.Sleep(time.Second)
	fmt.Println("done")

	// when we run this program, we see the output of the blocking call first, then the output of the two goroutines
	// the gouroutines' output may be interleaved, because goroutines are being run concurrently by the go runtime
}
