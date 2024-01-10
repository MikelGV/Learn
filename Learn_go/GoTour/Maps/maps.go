package main

import (
    "strings"
    "golang.org./x/tour/wc"
)

func WordCount(s string) map[string]int {
    maP := make(map[string]int)
    fields := strings.Fields(s)

    for _, i := range fields {
        maP[i]++
    }
    return maP
}

func main() {
    wc.Test(WordCount)
}
