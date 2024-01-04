package main

// This syntax imports the encoding/base64 package with the b64 name instead of
// the default base64
import (
    "fmt"
    b64 "encoding/base64"
)

func main() {
    // String we'll encode/decode
    data := "abc123!?$*&()'-=@~"

    // Support of both Standard and URL-compatible base64.
    // Encoder requires a []byte so we convert our string to that type
    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)

    // Decoding may return an error, which you can check if you don't already know
    // the input to be well-formed.
    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))
    fmt.Println()

    // This encodes/decodes using URL-compatible base64 format.
    uEnc := b64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc)
    uDec, _ := b64.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec))
}
