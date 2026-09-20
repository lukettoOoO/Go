// sometimes we'll want to sort a collection by something other than its natural order
// for example, suppose we wanted to sort strings by their length instead of alphabetically
package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	// we implement a comparison function for string lengths
	// cmp.Compare is helpful for this
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	// now we call slices.SortFunc with this custom comparison function to sort fruits by name length
	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	// we can use the same technique to sort a slice of values that aren't built-in types
	type Person struct {
		name string
		age int
	}

	people := []Person{
		Person{name: "Jax", age: 37},
		Person{name: "TJ", age: 25},
		Person{name: "Alex", age: 72},
	}

	// sort people by age using slics.SortFunc
	slices.SortFunc(people,
		func(a, b Person) int {
			// note: if the Person struct is large, you may want the slice to contain *Person instead and adjust the sorting function accordingly
			return cmp.Compare(a.age, b.age)
		})
	fmt.Println(people)
}
