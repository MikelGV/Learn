package main

import (
	"fmt"
	"time"
)

func main() {
    // Timer == event in the future
    timer1 := time.NewTimer(2 * time.Second)

    <-timer1.C
    fmt.Println("Timer 1 fired")

    // New timer that stops before fired
    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 fired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stoped")
    }

    time.Sleep(2 * time.Second)
}
