package main

import(
	"fmt"
	"unsafe"
)

func main() {
	var a int = 89
	b := 95
	fmt.Println("value of a is", a, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d bytes", a, unsafe.Sizeof(a)) // type and size of a
	fmt.Printf("\ntype of b is %T, size of b is %d bytes", b, unsafe.Sizeof(b)) // tppe and size of b
}
