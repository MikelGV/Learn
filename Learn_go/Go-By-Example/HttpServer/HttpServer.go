// Basic http server using the net/http package
package main

import (
	"fmt"
	"net/http"
)

// A fundamental concept in net/http servers is handlers. A handlers is an obj
// implementing the http.Handler interface. A common way to write a handler is
// by using the http.HandlerFunc adapter on functions with the apropiate signature.
func hello(w http.ResponseWriter, req *http.Request) {
    // Functions serving as handlers take a http.ResponseWriter and http.Request
    // as arguments. The response writer is used to fill in the HTTP response.
    // Here our simple response is just "hello\n"
    fmt.Fprintf(w, "hello\n")
}

// This handler does something a little more sophisticated by reading all HTTP
// request headers and echoing them into the response body.
func headers(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {
    // We register our handlers on server routes using the http.HandleFunc
    // convenience function. It sets up the default router in the net/http
    // package and takes a function as an argument.
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)

    // Finally, we call the ListAndServer with the port and a handler.nil tells
    // it to use the default router we've just set up.
    http.ListenAndServe(":8080", nil)
}
