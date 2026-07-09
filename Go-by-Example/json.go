// go offers built-in support for json encoding and decoding, including to and from built-in and custom data
package main

import (
	"encoding/json"
	"fmt"
	//"os"
	//"strings"
)

// we'll use these two structs to demonstrate encoding and decoding of custom types below
type response1 struct {
	Page int
	Fruits []string
}

// only exported fields will be encoded/decoded in json
// fields must start with capital letters to be exported
type response3 struct {
	Page int `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// first we'll look at encoding basic data types to json strings
	// here are some example for atomic values

	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// and here are some for slices and maps, which encode to json arrays and objects as you'd expect
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// the json package can automatically encode your custom data types
	// it will only include exported fields in the encoded output and will by default use those names as the json keys
	res1D := &response1{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	
}
