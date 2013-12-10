package main

import (
    "fmt"
)

func main() {
    var c chan
    fmt.Println(keyChannel)
    keyChannel(&c)
    for {
        key <- c
        fmt.Println(string(key))
    }
}

