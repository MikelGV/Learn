package main

import (
	"fmt"
	"time"
)

// Getting the number of seconds, milliseconds and nanoseconds from the unix Epoch
func main() {

    // Use time.Now with Unix, UnixMilli or UnixNano to get elapsed time since
    // the Unix epoch in seconds, milliseconds or nanoseconds, respectively.
    now := time.Now()
    fmt.Println(now)

    fmt.Println(now.Unix())
    fmt.Println(now.UnixMilli())
    fmt.Println(now.UnixNano())

    // Can also convert integer seconds or nanoseconds since the epoch into 
    // corresponding time.
    fmt.Println(time.Unix(now.Unix(), 0))
    fmt.Println(time.Unix(0, now.UnixNano()))
}
