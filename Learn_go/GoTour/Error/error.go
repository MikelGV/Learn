package main

import "fmt"

type ErrNegativeSqrt float64

func (err ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("Sqrt negative number: %f", float64(err))
}

func Sqrt(x float64) (float64, error) {
    if x < 0 {
        return  0, ErrNegativeSqrt(x)
    }

    z := x / 2

    for i := 0; i <= 1000; i++ {
        zPrev := z
        z -= (z*z - x) / (2*z)
        if z == zPrev {
            break
        }
    }
    return z, nil
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
