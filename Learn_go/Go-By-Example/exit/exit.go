package main
// Use os.Exit() to immediately exit with a given status.

import (
	"fmt"
	"os"
)

func main() {
    // Defers will not be run when using os.Exit, so this println will never be
    // called.
    defer fmt.Println("!")

    // Exit with status 3.
    os.Exit(3)
}
