// for more complex states we can use a mutex to safely access data across multiple goroutines
package main

import (
	"fmt"
	"sync"
)

// container hold a map of counters; since we want to update it concurrently from multiple goroutines, we add a Mutex to synchronize access
// note that mutexes must not be copied, so if this struct is passed around, it should be done by pointer
type Container struct {
	mu sync.Mutex
	counters map[string]int
}

// lock the mutex before accessing counters; unlock it at the end of the function using a defer statement
func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

// note the zero value of a mutex is usable as-is, so no initialization is required above
func main() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// this function increments a named counter in a loop
	doIncrement := func(name string, n int) {
		for range n {
			c.inc(name)
		}
	}

	// run several goroutines concurrently; note that they all accesss the same Container, and two of them access the same counter
	wg.Go(func() {
		doIncrement("a", 10000)
	})

	wg.Go(func() {
		doIncrement("a", 10000)
	})

	wg.Go(func() {
		doIncrement("b", 10000)
	})

	// wait for the goroutines to finish
	wg.Wait()
	fmt.Println(c.counters)

	// running the program shows that the counters updated as expected
}
