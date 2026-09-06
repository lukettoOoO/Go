package main

import (
	"fmt"
)

func main() {
	var size int
	fmt.Scanf("%d", &size);
	
	a := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &a[i])
	}
	
	//fmt.Println(a)
	max_len := 0

	for i := 0; i < size; i++ {
		for j := i; j < size; j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += a[k]
			}
			if sum == 0 {
				len := j - i + 1
				if len > max_len {
					max_len = len
				}
			}
		}
	}

	fmt.Println(max_len)
}
