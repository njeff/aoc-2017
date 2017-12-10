package main

import (
    "fmt"
    //"math"
)

func main() {
    //var input float64 = 289326
    a := make([][]int, 10)
    for i:=0; i<10; i++ {
        a[i] = make([]int, 10)
    }
    fmt.Print(a)
}