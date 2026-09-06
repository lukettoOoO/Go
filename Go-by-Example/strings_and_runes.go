// a go string is a read-only slice of bytes
// the language and the standard library treat strings specifically - as containers of text encoded UTF-8
// in other languages, strings are made of "characters"
// in go, the concept of a character is called a rune - it's an integer that represents a Unicode code point

package main

import (
	"fmt"
	"unicode/utf8"
)

func main () {
	// s is a string assigned a literal value representing the word "hello" in the Thai language
	// go string literals are utf-8 encoded text
	const s = "สวัสดี"

	// since strings are equivalent to []byte, this will produce the length of the raw bytes stored within
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// to count how many runes are in a string, we can use the utf8 package
	// note that the run-time of RuneCountInString depends on the size of the string, because it has to decode each utf-8 rune sequentially
	// some thai characters are represented by utf-8 code points that can be span multiple bytes, so the result of this count may be surprising
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// a range loop handles strings specially and decodes each rune along with its offset in the string
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// we can achieve the same iteration by using the utf8.DecodeRuneInString function explicitly
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		// this demostrates passing a rune value to a function
		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	// values enclosed in single quotes are rune literals
	// we can compare a rune value to a rune literal directly
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
