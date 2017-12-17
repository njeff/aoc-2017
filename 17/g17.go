package main

import (
    "fmt"
)

func main() {
    input := 312
    buffer := make([]int, 2018)
    buffer[0] = 0
    position := 0

    for i:= 1; i<=2017; i++ {
        position = (position+input)%i+1
        shift(buffer,position,i)
        buffer[position] = i
    }

    fmt.Println(buffer[position+1])
}

func shift(buf []int, pos int, l int) {
    for i:= l-1; i>pos; i-- {
        buf[i] = buf[i-1]
    }
}
