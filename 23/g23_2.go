package main

import (
    "fmt"
    "math"
)

func main() {
    b := 109900
    c := 126900

    out := 0
    for i := b; i<= c; i+= 17 {
        if !isPrime(i) {
            out++
        }
    }
    fmt.Println(out)
}


func isPrime(v int) bool {
    for i := 2; i<= int(math.Sqrt(float64(v))); i++ {
        if v%i == 0 {
            return false
        }
    }
    return true
}
