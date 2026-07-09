package main

import "fmt"

func main() {
	// Example:
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Println("Println:")
	fmt.Println(sample)

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("\n")

	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sample)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample)

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample)

	// Exercise: Modify the examples above to use a slice of bytes instead of a string. Hint: Use a conversion to create the slice
	fmt.Printf("\n\n")

	var byte_sample []byte;
	byte_sample = make([]byte, len(sample)) // or use byte_sample := []byte(sample)
	for i := 0; i < len(sample); i++ {
		byte_sample[i] = byte(sample[i])
		fmt.Printf("%x ", byte_sample[i])
	}
	fmt.Printf("\n")

	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", byte_sample)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", byte_sample)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", byte_sample)

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", byte_sample)

	// Exercise: Loop over the string using the %q format on each byte. What does the output tell you?
	fmt.Printf("\n\n")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%q", sample[i]) // indexing a string yields its bytes, not its characters
	}
}
