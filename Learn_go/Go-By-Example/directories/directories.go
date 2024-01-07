package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    // Create a new Sub-directory in the current directory.
    err := os.Mkdir("subdir", 0755)
    check(err)

    // When creating temporary directories, it's good practice to defer their
    // removal. Os.Removal will delete a whole directory tree (similarly to rm -rf).
    defer os.RemoveAll("subdir")

    // Helper function to create a new empty file.
    createEmptyFile := func(name string) {
        d := []byte("")
        check(os.WriteFile(name, d, 0644))
    }

    createEmptyFile("subdir/file1")

    // Can create a hierarchy of directories, including parents with MkdirAll.
    // This is similar to the command-line mkdir -p.
    err = os.MkdirAll("subdir/parent/child", 0755)
    check(err)

    createEmptyFile("subdir/parent/file2")
    createEmptyFile("subdir/parent/file3")
    createEmptyFile("subdir/parent/child/file4")

    // ReadDir lists directory contens, returning a slice of os.DirEntry objs.
    c, err := os.ReadDir("subdir/parent")
    check(err)

    fmt.Println("Listing subdir/parent")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

    // Chdir lets us change current working directory, similarly to cd.
    err = os.Chdir("subdir/parent/child")
    check(err)

    // Now we'll see the content of subdir/parent/child when listing the current
    // directory.
    c, err = os.ReadDir(".")
    check(err)

    fmt.Println("Listing subdir/parent/child")
    for _, entry :=range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

    // Cd back to where we started.
    err = os.Chdir("../../..")
    check(err)

    // We can also visit a directory recursively, including all its sub-directories
    // WalkDir accepts a callback function to handle every file or directory visited.
    fmt.Println("Visiting subdir")   
    err = filepath.WalkDir("subdir", visit)
}

// visit Is called for every file or directory found recursively by filepath.WalkDir.
func visit(path string, d fs.DirEntry, err error) error {
    if err != nil {
        return err
    }

    fmt.Println(" ", path, d.IsDir())
    return nil
}