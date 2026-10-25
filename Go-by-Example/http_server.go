// writing a basic http server is easy using the net/http package
package main

import (
	"fmt"
	"net/http"
)

// a fundamental concept in net/http servers is handlers
// a handler is an object implementing the http.Handler interface
// a common way to write a handler is by using the http.HandlerFunc adapter on functions with the appropriate signature

func hello(w http.ResponseWriter, req *http.Request) {
	// functions serving as handlers take a http.ResponseWriter and a http.Request as arguments
	// the response writer is used to fill in the http response
	// here the response is just "hello\n"
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	// this handler does something a little more sophisticated by reading all the http request headers and echoing them into the response body
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	// we register our handlers on server routes using the http.HandleFunc convenience function
	// it sets up the default router in the net/http package and takes a function as an argument
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// finally, we call the ListenAndServe with the port and a handler
	// nil tells it to use the default router we've just set up
	http.ListenAndServe(":8090", nil)
}

// run the server in the background and access the /hello and /headers route
