package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
    // SHA256 hashes are frequently used to compute short identities for binary
    // or text blobs. Example, TLS/SSL certificates use SHA256 to compute a
    // certificate's signature.

    // Implements several hash functions in various crypto/*.

    // Here whe start with a new hash.
    s := "sha256 this string"

    h := sha256.New()

    // Write expects bytes. If string s, use []byte(s) to coerce it to bytes.
    h.Write([]byte(s))

    // This gets the finalized hash result as byte slice. Argument to sum can be
    // used to append to an existing byte slice: it usually isn't needed.
    bs := h.Sum(nil)


    fmt.Println(s)
    fmt.Printf("%x\n", bs)
}
