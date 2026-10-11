// sha256 hashes are frequently used to compute short identities for binary or text blobs
// for example, tls/ssl certificates use sha256 to compute a certificate's signature
// here's how to compute sha256 hashes in go
package main

// go implements severak hash functions in various crypto/* packages
import(
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"

	// here we start with a new hash
	h := sha256.New()

	// Write expects bytes
	// if we have a string s, use []byte(s) to coerce it to bytes
	h.Write([]byte(s))

	// this gets the finalized hash result as a byte slice
	// the argument to Sum can be used to append to an existing byte slice: it usually isn't needed
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)

	// running the program computes the hash and prints it in a human-readable hex format

	// we can compute other hashes using a similar pattern to the one shown above
	// for example, to compute sha512 hashes import crypto/sha512 and use sha512.New()
	// note that if we need cryptopgraphically secure hashes, we should carefully research hash strength
}
