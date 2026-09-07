// tickers are for when you want to do something repeatedly at regular intervals
package main

import (
	"time"
	"fmt"
)

func main() {
	// tickers are similar to timers: a channel that is sent values
	// here we'll use the select builtin on the channel to await the values as they arrive every 500ms
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <- done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// tickers can be stopped like timers
	// once a ticker is stopped it won't receive any more values on its channel
	// we'll stop ours after 1600ms
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")

	// when we run this program the ticker should tick 3 times before we stop it
}
