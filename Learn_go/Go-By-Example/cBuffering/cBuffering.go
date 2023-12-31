package main

import "fmt"

func main() {
    fmt.Println("this is channels buffering")
    message := make(chan string, 2)

    message <- "buffering"
    message <- "channel"

    fmt.Println(<-message)
    fmt.Println(<-message)
}

