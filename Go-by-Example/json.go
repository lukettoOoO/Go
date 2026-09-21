// go offers built-in support for json encoding and decoding, including to and from built-in and custom data
package main

import (
	"bytes"
	"encoding/json/v2"
	"fmt"
	"strings"
)

// we'll use these two structs to demonstrate encoding and decoding of custom types below
type response1 struct {
	Page int
	Fruits []string
}

// only exported fields will be encoded/decoded in json
// fields must start with capital letters to be exported
type response2 struct {
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

	// we can use tags on struct field declarations to customize the encoded json key names
	// check the definition of response2 above to see an example of such tags
	res2D := &response2{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// now let's look at decoding json data into go values
	// here's an example for a generic data structure
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// we need to provide a variable where the json package can put the decoded data
	// this map[string]any will hold a map os strings to arbitrary data types
	var dat map[string]any

	// here's the actual decoding, and a check for associated errors
	// for the sake of brevity we ignore the errors in these examples; in real code, we should always check for errors and act upon them
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// in order to use the values in the decoded map, we'll need to convert them to their appropriate type
	// for example here we convert the value in num to the expected float64 type
	num := dat["num"].(float64)
	fmt.Println(num)

	// accessing nested data requires a series of conversions
	strs := dat["strs"].([]any)
	str1 := strs[0].(string)
	fmt.Println(str1)

	// we can also decode json into custom data types
	// this has the advantages of adding additional type-safety to our programs and eliminating the need for type assertions when accessing the decoded data
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	_ = json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// in the examples above we always used bytes and strings as intermediates between the data and json representation on standard out
	// we can also stream json encodings directly to io.Writers like os.Stdout or even http respons bodies
	d := map[string]int{"apple": 5, "lettuce": 7}
	var buf bytes.Buffer
	_ = json.MarshalWrite(&buf, d)
	fmt.Println(buf.String())

	// streaming reads from io.Readers like os.Stdin or http request bodies is done with json.UnmarshalRead
	res1 := response2{}
	_ = json.UnmarshalRead(strings.NewReader(str), &res1)
	fmt.Println(res1)
}
