package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    // slice of dy
    s := make([][]uint8, dy)
    for i := 0; i < dy; i++ {
        //slice of dx
        ss := make([]uint8, dx)
        for x := 0; x < dx; x++ {
            ss[x] = uint8(x ^ i)
        }
        s[i] = ss
    }
    return s
}

func main() {
    pic.Show(Pic)
}
