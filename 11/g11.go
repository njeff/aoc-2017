package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "math"
)

func main() {
    b, err := ioutil.ReadFile("input.txt")
    if err != nil {
        fmt.Print(err)
    }
    l := strings.Split(strings.TrimSpace(string(b)),",")

    x, y, z := 0.0, 0.0, 0.0
    max := 0.0

    for _, e := range l {
        c := string(e)
        switch c {
        case "n":
            y++
            z--
        case "s":
            y--
            z++
        case "ne":
            x++
            z--
        case "nw":
            x--
            y++
        case "se":
            x++
            y--
        case "sw":
            x--
            z++
        }
        td := math.Max(math.Abs(x),math.Max(math.Abs(y),math.Abs(z)))
        if td > max {
            max = td
        }
    }
    d := math.Max(math.Abs(x),math.Max(math.Abs(y),math.Abs(z)))
    fmt.Println(d)
    fmt.Println(max)
}