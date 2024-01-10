package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr
func (i IPAddr) String() string {
    s := make([]string, len(i))
    for x, val := range i {
        s[x] = strconv.Itoa(int(val))
    }
    return fmt.Sprintf("%d.%d.%d.%d", i[0], i[1], i[2], i[3])
}

func main() {
    hosts := map[string]IPAddr{
        "loopback": {127, 0, 0, 1},
        "googleDNS": {8, 8, 8, 8},
    }
    for name, ip := range hosts {
        fmt.Printf("%v: %v\n", name, ip)
    }
}
