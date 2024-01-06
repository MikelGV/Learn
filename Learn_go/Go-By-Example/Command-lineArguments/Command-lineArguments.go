// Comman-line Arguments are a common way to parameterize execution of programs.
// For example, go run hello.go uses a run and hello.go arguments to the go pg.
package main

import (
	"fmt"
	"os"
)

func main () {
    // Os.Args provides access to raw comman-line arguments.
    // Note that first value in this slice is the path to the program, and
    // os.Args[1:] holds the arguments to the program.
    argsWithProg := os.Args
    argsWihtoutProg := os.Args[1:]

    // Can get individual args with normal indexing.
    arg := os.Args[3]

    fmt.Println(argsWithProg)
    fmt.Println(argsWihtoutProg)
    fmt.Println(arg)
}
