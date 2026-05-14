package main

import "fmt"

func main() {
	s := []int{2, 3, 6, 7, 11, 13}
	s = s[1:4]
	fmt.Println(s) // 3, 6, 7
	
	s = s[:2] // 3, 6
	fmt.Println(s)

	s = s[1:] // 6
	fmt.Println(s)
}
