// the go standard library comes with excellent support for http clients and servers in the net/http package
// in this example we'll use it to issue simple http requests
package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	// issue an http GET request to a server
	// http.Get is a convenient shortcut around creating an http.Client object and calling its Get method; it uses the http.DefaultClient object which has useful default settings
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// print the http response status
	fmt.Println("Response status:", resp.Status)

	// print the first 5 lines of the response body
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
