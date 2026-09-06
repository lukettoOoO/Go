package main

import "fmt"

func sumAll(list []int) int {
	sum := 0
	for i := 0; i < len(list); i++ {
		sum += list[i]
	}
	return sum
}

func main() {
    list := []int{1, 2, 3, 4, 5, 6}

    fmt.Println(sumAll(list))
}
