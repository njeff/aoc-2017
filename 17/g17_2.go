package main

import (
    "fmt"
)


func main() {
    input := 312
    position := 0
    zp := 0
    az := 0

    for i:= 1; i<=5E7; i++ {
        position = (position+input)%i+1
        if position == zp+1 {
            az = i
        }
        if position == zp {
            zp++
        }
    }
    fmt.Println(az)
}

