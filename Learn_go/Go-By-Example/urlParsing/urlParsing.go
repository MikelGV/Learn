package main

import (
	"fmt"
	"net"
	"net/url"
)

// URLs provide a uniform way to locate resources.
func main() {
    
    // Parse this example URL, which includes a scheme, authentication info, host
    // port, path, query params, and query fragment.
    s := "postgres://user:pass@host.com:5432/path?k=v#f"

    // Parse the URL and ensure there are no errors.
    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }

    // Accessing the scheme is straightforward.
    fmt.Println(u.Scheme)

    // User contains all authentication info; call Username and Password on this
    // For individual values.
    fmt.Println(u.User)
    fmt.Println(u.User.Username())
    p, _ := u.User.Password()
    fmt.Println(p)

    // The host contains both hostname and the port. Use SplitHostPort to extract.
    fmt.Println(u.Host)
    host, port, _ := net.SplitHostPort(u.Host)
    fmt.Println(host)
    fmt.Println(port)

    // Extract the path and the fragment after the #.
    fmt.Println(u.Path)
    fmt.Println(u.Fragment)

    // Get the query params in a string of k=v format, use RawQuery.
    // Can also parse query params into a map. The parsed query param maps are
    // from string to slices of strings, so index into [0] if you want the first value.
    fmt.Println(u.RawQuery)
    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m)
    fmt.Println(m["k"][0])
}
