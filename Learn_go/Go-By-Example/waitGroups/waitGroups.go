package main

import (
	"fmt"
	"sync"
	"time"
)

// To wait for multiple goroutines to finish, we can use a wait group

// this function will run in every goroutine
func worker(id int) {
    fmt.Printf("Worker %d starting\n", id)

    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}


func main() {
    fmt.Println("this is wait groups")

    // This wait group is used to wait for all goroutines launched here to finish
    var wg sync.WaitGroup

    // Launch several go routines and increment wait groups counter for each
    for i := 1; i <= 5; i++ {
        wg.Add(1)
       
        // Avoid re-use of the same i value in each goroutine closure
        i := i

        // Wrap the worker call in a closure that makes sure to tell waitGroup
        // that this worker is done.
        go func() {
            defer wg.Done()
            worker(i)
        }()
    }

    // Block until the wait group counter goes back to 0
    wg.Wait()
}
