// How to use the net/http package for http client
package main

import (
	"bufio"
	"fmt"
	"net/http"
)


func main() {
    // Issue an Http.GET request to a server. http.GET is a convenient shortcut
    // around creating an http.Client obj and calling its get method; it uses the
    // http.DefaultClient obj which has useful default settings.
    resp, err := http.Get("https://gobyexample.com")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    // Print the Http response status.
    fmt.Println("Response status:", resp.Status)

    // Print the first 5 lines of the response body.
    scanner := bufio.NewScanner(resp.Body)
    for i := 0; scanner.Scan() && i < 5; i++ {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}
