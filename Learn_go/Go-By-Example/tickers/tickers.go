package main

import (
	"fmt"
	"time"
)

func main() {
    // Tickers are for when you want to do something
    // repeatedly at regular intervals

    // Ticker use a similar mechanism to timers
    ticker := time.NewTicker(500 * time.Millisecond)
    done := make(chan bool)

    go func(){
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                fmt.Println("Tick at:", t)
            }
        }
    }()

    // Tickers can be stopped like timers
    time.Sleep(1600 * time.Millisecond)
    ticker.Stop()
    done <- true
    fmt.Println("Ticker stopped")

}
