// Sometimes we'd like our Go programs to intelligently handle Unix signals.
// For example, we might want a server to gracefully shutdown when it receives a
// SIGTERM, or a SIGINT. Here's how to handle signals in Go with channels.
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
    // Go signal notification works by sending os.Signal values on a channel.
    // We'll create a channel to receive these notifications. Note that this
    // channel should be buffered.
    sigs := make(chan os.Signal, 1)

    // Signal.Notify registers the given channel to receive notifications of the
    // specified signals.
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    // We could receive from sigs here in the main function, but let's see how
    // this could also be done in a separate goroutine, to demostrate a more
    // realistic scenario of graceful shutdown.
    done := make(chan bool, 1)

    // This goroutine executes a blocking receive for signals. When it gets one
    // it'll print it out and then notify the program that it can finish.
    go func (){
        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)
        done <-true
    }()

    // The program will wait here until it gets the expected signal (as indicated
    // by the gorouting above sending a value on done) and then exit.
    fmt.Println("awaiting signal")
    <-done
    fmt.Println("exiting")
}
