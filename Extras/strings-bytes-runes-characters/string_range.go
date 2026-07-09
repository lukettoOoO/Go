package main

import "fmt"

func main() {
	const nihongo = "日本語"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}

	// Exercise: Put an invalid UTF-8 byte sequence into the string. (How?) What happens to the iterations of the loop?
	const invalid = "A\xbd🚀Å"
	for index, runeValue := range invalid {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	} // handles the situation gracefully
}
