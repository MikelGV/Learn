package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

    // This test whether a pattern matches a string
    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Println(match)

    // Compile and optimize regex struct
    r, _ := regexp.Compile("p([a-z]+)ch")

    // Many methods available on these structs. Here's a match test like we saw.
    fmt.Println(r.MatchString("peach"))

    // Finds the match for regexp.
    fmt.Println(r.FindString("peach punch"))

    // Finds the match but returns start and end indexes for the match
    fmt.Println("idx:", r.FindStringIndex("peach punch"))

    // Submatch variants include info about both the whole-pattern matches and
    // sub-matches within those matches
    fmt.Println(r.FindStringSubmatch("peach punch"))

    // Find info about the indexes of matches and sub-matches
    fmt.Println(r.FindStringSubmatchIndex("peach punch"))

    // Find All variants of these functions apply to all matches in input
    fmt.Println(r.FindAllString("peach punch pinch", -1))

    // All variants available for the other functions we saw above
    fmt.Println("all:", r.FindAllStringSubmatchIndex("peach punch pinch", -1))

    // non-negative integer as the second argument limit number of matches
    fmt.Println(r.FindAllString("peach punch pinch", 2))

    // examples above had string args and used names like MatchString. We can
    // provide []byte args and drop string from func name
    fmt.Println(r.Match([]byte("peach")))

    // When creating global vars with regex can use the MustCompile variation
    // of Compile. MustCompile panics instead of returning an err which makes it
    // safer to use fro global variables.
    r = regexp.MustCompile("p([a-z]+)ch")
    fmt.Println("regexp:", r)

    // Regexp package can be used to replace subsets of string with other values
    fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

    // The Func variant allows to transform matched text with a given function
    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))
}
