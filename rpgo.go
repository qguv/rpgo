package main

import (
    "fmt"
)

func main() {
    var c chan
    go keyChannel(&c)
    for {
        key <- c
        fmt.Println(string(key))
    }
}
