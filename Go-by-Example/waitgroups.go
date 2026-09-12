// to wait for multiple goroutines to finish, we can use a wait group
package main

import (
	"fmt"
	"sync"
	"time"
)

// this is the function we'll run in every goroutine
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	// sleep to simulate an expensive task
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// this WaitGroup is used to wait for all the goroutines launched here to finish
	// note: if a WaitGroup is explicitly passed into functions, it should be done by pointer
	var wg sync.WaitGroup

	// lauch several goroutines using WaitGroup.Go
	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			worker(i)
		})
	}

	// block until all goroutines started by wg are done
	// a gorutine is done when the function it invokes returns
	wg.Wait()

	// note that this approach has no straightforward way to propagate errors from workers
	// for more advanced use cases, consider using the errgroup package

	// the order of workers starting up and finishing is likely to be different for each invocation
}
