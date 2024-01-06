package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Throughout program execution, we often want to create data that isn't needed
// after  program exits. Temprorary files and directories are useful for this
// purpose since they don't pollut the fs over time.

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// Easiest way to create a temporary file is by calling os.CreateTemp.
// Creates a file and opens it for reading and writing.
// Provide "" as the first arg, so os.CreateTemp will create the file in the
// defaul location for our OS.
func main() {
    f, err := os.CreateTemp("", "sample")
    check(err)

    // Display the name of the temporary file.
    fmt.Println("Temp file name:", f.Name())

    // Clean up the file after we're done.
    defer os.Remove(f.Name())

    // Write some data to file.
    _, err = f.Write([]byte{1,2,3,4})
    check(err)

    // If we intend to write many temporary files, we may prefer to create a
    // temporary directory. os.MkdirTemp arg are the same as CreateTemp, but
    // it returns a directory name rather than an open file.
    dname, err := os.MkdirTemp("", "sampledir")
    check(err)
    fmt.Println("Temp dir name:", dname)

    defer os.RemoveAll(dname)

    // Now we can synthesize temporary file names by prefixing them with our
    // temporary directory.
    fname := filepath.Join(dname, "file1")
    err = os.WriteFile(fname, []byte{1, 2}, 0666)
    check(err)
}
