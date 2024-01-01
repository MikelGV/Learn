package main

import (
	"fmt"
	"time"
)

func main() {
    fmt.Println("this is rate limiting")

    // limit incoming requests
    request := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        request <- i
    }
    close(request)

    // limiter channel will receive value every 200 ms
    limiter := time.Tick(200 * time.Millisecond)

    // blocking on a receive from the limiter before serving each requests
    // limit ourselfs to 1 req every 200 ms
    for req := range request {
        <-limiter
        fmt.Println("request", req, time.Now())
    }

    // allow short burst of requests in our rate limiting scheme preserving 
    // overall rate limit. Accomplish this by buffering limiter channel
    burstyLimiter := make(chan time.Time, 3)

    // fill up channel to represent allowed burst
    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
    }

    // every 200 ms try add new value to burstyLimiter up to limit
    go func() {
        for t := range time.Tick(200 * time.Millisecond) {
            burstyLimiter <- t
        }
    }()

    // simulate 5 more incoming requests. First three benefit from burst capability
    burstyRequest := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        burstyRequest <- i
    }
    close(burstyRequest)
    
    for req := range burstyRequest {
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
}
