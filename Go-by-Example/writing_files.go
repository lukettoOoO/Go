// writing files in go follows similar pattterns to the ones we saw earlier for reading

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// to start, here's how to dump a string (or just bytes) into a file
	d1 := []byte("hello\ngo\n")
	path1 := filepath.Join( "dat1")
	err := os.WriteFile(path1, d1, 0644)
	check(err)

	// for more granular writes, open a file for writing
	path2 := filepath.Join("dat2")
	f, err := os.Create(path2)
	check(err)

	// it's idiomatic to defer a Close immediately after opening a file
	defer f.Close()

	// you can Write byte slices as you'd expect
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// a WriteString is also available
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// issue a Sync to flush writes a stable storage
	f.Sync()

	// bufio provides buffered writers in addition to the buffered readers we saw earlier
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	// use Flush to ensure all buffered operations have been applied to the underlying writer
	w.Flush()
}
