package main

// go:embed is a compailer directive that allows programs to include arbitrary
// files and folders in the Go binary at build time.

import (
    "embed"
)

// Embed directives accept paths relative to the directory containing the Go
// source file. This directive embeds the contents of the file into string var
// immediately following it.

//go:embed folder/single_file.txt
var fileString string

// Or embed the contents of the file into a []byte.
//go:embed folder/single_file.txt
var fileByte []byte

// We can also embed multiple files or folders with wildcards. This uses vars
// of the embed.Fs type, which implments a simple virtual file system.

//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {
    // Print out the contents of single_file.txt
    print(fileString)
    print(string(fileByte))

    // Retrives some files from the embedded folder
    content1, _ := folder.ReadFile("folder/file1.hash")
    print(string(content1))

    content2, _ := folder.ReadFile("folder/file2.hash")
    print(string(content2))
}
