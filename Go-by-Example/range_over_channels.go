// we can use for and range ro iterate over values received from a channel
package main

import "fmt"

func main() {
	// we'll iterate over 2 values in the queue channel
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// this range iterates over each element as it's received from queue
	// because we closed the channel above, the iteration terminates after receiving the 2 elements
	for elem := range queue {
		fmt.Println(elem)
	}

	// this example also showed that it's possible to close a non-empty channel but still habe the remaining values be received
}
