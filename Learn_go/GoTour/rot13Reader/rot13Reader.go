package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
    r io.Reader
}

func rot13R(b byte) byte {
    switch {
    case b >= 65 && b <= 77:
        fallthrough
    case b >= 97 && b <= 109:
        b = b + 13
    case b >= 78 && b <= 90:
        fallthrough
    case b >= 110 && b <= 122:
        b = b - 13
    }
    return b
}

func (r *rot13Reader) Read(b []byte) (int, error) {
    n, err := r.r.Read(b)
    for i := 0; i <= n; i++ {
        b[i] = rot13R(b[i])
    }
    return n, err
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
