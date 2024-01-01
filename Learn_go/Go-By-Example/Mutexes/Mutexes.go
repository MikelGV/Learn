package main

import (
	"fmt"
	"sync"
)

// Container holds a map of counters
type Container struct {
    mu  sync.Mutex
    counters map[string]int
}

// Lock the mutex before accessing counters; unlock it at the end of the func
func (c *Container) inc(name string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.counters[name]++
}

// Note the zero value of a mutex is usable as-is, no initialization needed
func main() {
    fmt.Println("this is how mutexes work")

    c := Container{

        counters: map[string]int{"a": 0, "b": 0},
    }

    var wg sync.WaitGroup

    // This function increments a named counter in a loop
    doIncrement := func(name string, n int) {
        for i := 0; i < n; i++ {
            c.inc(name)
        }
        wg.Done()
    }

    // Run several goroutines concurrently;
    wg.Add(3)
    go doIncrement("a", 10000)
    go doIncrement("a", 10000)
    go doIncrement("b", 10000)

    // Wait for the goroutines to end
    wg.Wait()
    fmt.Println(c.counters)
}
