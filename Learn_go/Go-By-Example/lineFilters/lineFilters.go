package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// A line filter is a common type of program that reads input on stdin,
// and then processes it and then prints some derived result to stdout.grep and
// sed are common line filters.
func main() {
    
    // Wrapping the unbuffered os.Stdin with the buffered scanner gives us a,
    // convenient Scan method that advances the scanner to the next token.
    scanner := bufio.NewScanner(os.Stdin)

    // Text returns the current token from the input.
    for scanner.Scan() {

        // Write out the uppercased line.
        ucl := strings.ToUpper(scanner.Text())

        fmt.Println(ucl)
    }

    // Check for errors during Scan. End of file is expected and not reported
    // by scan as an error
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error", err)
        os.Exit(1)
    }
}
