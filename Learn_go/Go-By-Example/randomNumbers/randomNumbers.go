package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Go's math/rand package provides pseudorandom number generation.
func main() {

    // Rand.Intn returns a random int n,0 <= n < 100
    fmt.Print(rand.Intn(100), ",")
    fmt.Print(rand.Intn(100))
    fmt.Println()

    // Rand.Float64 returns a float64 f,0.0 <= f < 1.0.
    fmt.Println(rand.Float64())

    // This can be used to generate random floats in other ranges.
    fmt.Print((rand.Float64()*5)+5, ",")
    fmt.Print((rand.Float64() * 5) + 5)
    fmt.Println()

    // The default number generator is deterministic producing the same sequence
    // of numbers each time by default. Producing varying sequences, give it a
    // seed that changes. It's not safe to use random numbers you intend to be
    // secret; use crypto/rand for those.
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

    // Call the resulting rand.Rand just like the functionso on the rand package.
    fmt.Print(r1.Intn(100), ",")
    fmt.Print(r1.Intn(100))
    fmt.Println()

    // Seeding a source with the same number produces the same sequence of
    // random numbers.
    s2 := rand.NewSource(42)
    r2 := rand.New(s2)
    fmt.Print(r2.Intn(100), ",")
    fmt.Print(r2.Intn(100))
    fmt.Println()

    s3 := rand.NewSource(42)
    r3 := rand.New(s3)
    fmt.Print(r3.Intn(100), ",")
    fmt.Print(r3.Intn(100))
    fmt.Println()
}
