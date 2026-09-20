// go's slices package implements sorting for builtins and user-defined types
package main

import (
	"fmt"
	"slices"
)

func main() {
	// sorting functions are generic, and work for any ordered built-in type
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	// an example of sorting ints
	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:   ", ints)

	// we can also use the slices package to check if a slice is already in sorted order
	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)
}
