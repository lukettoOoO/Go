package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				goto END
			}
			fmt.Println(i, j)
		}
	}
END:
}
