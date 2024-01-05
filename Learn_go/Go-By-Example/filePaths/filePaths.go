package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// The filepath package provides functions to parse and construct file paths
// in a way that is portable between op systems; dir/file on linux vs dir\file
// on windows
func main() {

    // Join should be used to construct paths in a portable way. Takes any
    // number of args and constructs a hierarchical path for them.
    p := filepath.Join("dir1", "dir2", "filename")
    fmt.Println("p:", p)

    // Should allways use Join instead of concatenating /s or \s manually.
    // In addition to providing portability, Join will normalize paths by
    // removing superfluous separators and change directory changes.
    fmt.Println(filepath.Join("dir1//", "filename"))
    fmt.Println(filepath.Join("dir1/../dir1", "filename"))

    // Dir and Base can be used to split a path to the directory and file.
    // Alternatively, Split will return both in the same call.
    fmt.Println("Dir(p):", filepath.Dir(p))
    fmt.Println("Base(p):", filepath.Base(p))

    // Can check if path is absolute.
    fmt.Println(filepath.IsAbs("dir/file"))
    fmt.Println(filepath.IsAbs("/dir/file"))

    filename := "config.json"

    // Some file names have extensions following dot. We can split the extension
    // out of such names with Ext.
    ext := filepath.Ext(filename)
    fmt.Println(ext)

    // Find file's name with the extension removed, use strings.TrimSuffix.
    fmt.Println(strings.TrimSuffix(filename, ext))

    // Rel finds a relative path between a base and a target. It returns an err
    // if the target can't be made relative to base.
    rel, err := filepath.Rel("a/b", "a/c/t/file")
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)
}
