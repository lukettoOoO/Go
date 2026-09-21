// go offers built-in support for regular expressions
// here are some examples of common regexp-related tasks in go
package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	// this tests whether a pattern matches a string
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// above we used a string pattern directly, but for other regexp tasks we'll ned to Compile an optimized Regexp struct
	r, _ := regexp.Compile("p([a-z]+)ch")

	// many methods are available on these structs
	// here's a match test like we saw earlier
	fmt.Println(r.MatchString("peach"))

	// this finds the match for the regexp
	fmt.Println(r.FindString("peach punch"))

	// this also finds the first match but returns the start and end indexes for the match instead of the matching text
	fmt.Println("idx:", r.FindStringIndex("peach punch"))

	// the Submatch variants include information about both the whole-pattern matches and the submatches within those matches
	// for example this will return information for both p([a-z]+)ch and ([a-z]+)
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// similarly this will return information about the indexes of matches and submatches
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// the All variants of these functions apply to all matches in the input, not just the first
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// these All variants are available for the other functions we saw above as well
	fmt.Println("all:", r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	// providing a non-negative integer as the second argument of these functions will limit the number of matches
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// our examples above had string arguments and used names like MatchString
	// we can also provide []byte arguments and drop String from the function name
	fmt.Println(r.Match([]byte("peach")))

	// when creating global variables with regular expressions we can use the MustCompile variation of Compile
	// MustCompile panics instead of returning an error, which makes it safer to use for global variables
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	// the regexp package can also be used to replace subsets of strings with other values
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// the Func variant allows you to transform matched text with a given function
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}

