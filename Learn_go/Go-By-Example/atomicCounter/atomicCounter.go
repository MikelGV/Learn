package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
    // Let's use sync/atomic package to manage state over communication of channels
    fmt.Println("this is atomic counter")

    // atomic integer to represent (always positive) counter
    var ops atomic.Uint64

    // waitgroup wait for all goroutines to finish their work
    var wg sync.WaitGroup

    // start with 50 goroutines increment counter 100 times
    for i := 0; i < 50; i++ {
        wg.Add(1)

        go func() {
            for c := 0; c < 100; c++ {
                
                // atomically increase counter use Add()
                ops.Add(1)
            }
            wg.Done()
        }()
    }

    // wait till all routines are done
    wg.Wait()

    // no goroutines writing to ops but using load it's safe to atomically
    // read value even while other routines are (atomically) updating it
    fmt.Println("ops:", ops.Load())
}
