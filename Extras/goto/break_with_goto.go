package main

import "fmt"

func main() {
	for i := 0; ; i++ {
		if i == 10 {
			goto END
		}
		fmt.Print(i)
	}
END:
	fmt.Println("END")
}
