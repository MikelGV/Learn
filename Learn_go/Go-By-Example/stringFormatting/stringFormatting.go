package main

import (
	"fmt"
	"os"
)

type point struct {
    x, y int
}

func main() {
    // Go offers execellent support for string formatting in the printf tradition
    // here are some examples of common string formatting tasks.

    // This prints an instance of our point struct.
    p := point{1, 2}
    fmt.Printf("struct1: %v\n", p)

    // If the value is a struct, the %v variant will include the struct field name
    fmt.Printf("struct2: %+v\n", p)

    // The %#v variant prints a Go syntax representation of the value
    fmt.Printf("struct3: %#v\n", p)

    // To print the type of value use %T
    fmt.Printf("type: %T\n", p)

    // Formatting booleans is straight-forward
    fmt.Printf("bool: %t\n", true)
    
    // Use %d for standard, base-10 formatting
    fmt.Printf("int: %d\n", 123)
    
    // This prints a binary representation
    fmt.Printf("bin: %b\n", 14)

    // This prints the character corresponding to the given int
    fmt.Printf("char: %c\n", 33)

    // %x provides hex encoding
    fmt.Printf("hex: %x\n", 456)
 
    // For basic decimal formatting use %f
    fmt.Printf("float1: %f\n", 48.9)
    
    // %e and %E format the float in (slightly different version of) science notation
    fmt.Printf("float2: %e\n", 123400000.0)
    fmt.Printf("float3: %E\n", 123400000.0)
    
    // For basic sting printing use %s
    fmt.Printf("str1: %s\n", "\"string\"")

    // Double-coute strings as in Go source, use %q
    fmt.Printf("str2: %q\n", "\"string\"")

    // As with integers seen earlier %x renders the string in base-16
    fmt.Printf("str3: %x\n", "hex this")

    // To print a representation of a pointer use %p
    fmt.Printf("pointer: %p\n", &p)

    // To specify width of an integer use a number after %
    fmt.Printf("width: |%6d|%6d|\n", 12, 345)

    // You can also specify widht of printed floats
    fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

    // To left-justify use the - flag
    fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

    // Controll the width of strings, right-justified width
    fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

    // To left-justify use the - flag
    fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

    // os.Stdout.Sprintf formats and returns a string without printing it anywhere
    s := fmt.Sprintf("sprintf: a %s", "string")
    fmt.Println(s)

    // format+print to io.Writers other than os.Stdout using Fprintf
    fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
