package main

import "fmt"

func main() {
	var i = 0
BEGIN:
	fmt.Printf("%d ", i)
	if i == 9 {
		goto END
	}
	i++
	goto BEGIN
END:
}
