package main
// Program only takes about 2 seconds despite doing about 5 seconds of work

import (
	"fmt"
	"time"
)

// workers receive work on the jobs channel and send the corresponding results
func worker(id int, jobs <-chan  int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "started job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }
}

func main(){
    fmt.Println("this are worker pools")
    
    // Send work and collect results channels
    const numJobs = 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    // Starts 3 workers initialy blocked because there is no jobs yet
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Send 5 jobs and then close that channel
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    
    close(jobs)

    // Finally collect all the results ensuring that all goroutines are closed
    for a := 1; a <= numJobs; a++ {
        <-results
    }
}
