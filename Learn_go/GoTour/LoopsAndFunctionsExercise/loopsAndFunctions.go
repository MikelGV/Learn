package main

import "fmt"

func Sqrt(x float64) float64 {
    if x < 0 {
        return -1
    } else {
        return 0
    }

    z := x / 2

    for i := 0; i <= 1000; i++ {
        zPrev := z
        z -= (z*z - x) / (2*z)
        if z == zPrev {
            break
        }
    }
    return z
}

func main() {
    fmt.Println(Sqrt(2))
}
