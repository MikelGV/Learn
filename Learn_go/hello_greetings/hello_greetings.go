package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    // Set properties of the predefined logger.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // Get a greeting message and print it.
    //message := greetings.Hello("Mikel")
    //fmt.Println(message)

    // Request a greeting message.
    message, err := greetings.Hello("Mikel")
   // Look for non-nil error value
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(message)
}
