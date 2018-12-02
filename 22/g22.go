package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)

func main() {
    b, err := ioutil.ReadFile("input.txt")
    if err != nil {
        fmt.Print(err)
    }
    s := strings.TrimSpace(string(b))
    input := strings.Split(s, "\n")
    width := len(input[0])
    fmt.Println(width)

    g := [][]rune{}
    ls := ""
    for i := 0; i<1001; i++ {
        ls += "."
    }
    ed := ""
    for i := 0; i<(1001-width)/2; i++ {
        ed += "."
    }
    for i := 0; i<1001; i++ {
        g = append(g, []rune(ls))
    }
    for _, e := range input {
        t := []rune(ed + e + ed)
        g = append(g, t)
    }
    for i := 0; i<1000; i++ {
        g = append(g, []rune(ls))
    }
    px := len(g[500])/2
    py := len(g)/2

    direction := 0 //0 up, increasing cw

    inf := 0
    for i := 0; i<10000; i++ {
        if g[py][px] == '#' {
            direction = (direction+1)%4
            g[py][px] = '.'
        } else {
            direction = (direction+3)%4
            g[py][px] = '#'
            inf++
        }
        switch direction {
        case 0:
            py--
        case 1:
            px++
        case 2:
            py++
        case 3:
            px--
        }
    }
    fmt.Println(inf)
}


