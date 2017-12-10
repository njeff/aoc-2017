package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    b, err := ioutil.ReadFile("input.txt")
    if err != nil {
        fmt.Print(err)
    }
    s := string(b)
    array := []rune(s)
    total := 0
    for k := range array{
        if array[k] == array[(k+1)%len(array)] {
            total += int(array[k] - '0')
        }
    }
    fmt.Println(total)
}