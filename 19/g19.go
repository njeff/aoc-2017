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
    s := string(b)
    input := strings.Split(s, "\n")

    direction := 1 //1 = down, 2 = left, 3 = up, 4 = right
    px := 0
    py := 0
    for i, e := range input[0] {
        if e == '|' {
            px = i
            break
        }
    }
    output := ""
    steps := 0

    for input[py][px] != ' ' {
        steps++
        switch direction {
        case 1:
            py++
        case 2:
            px--
        case 3:
            py--
        case 4:
            px++
        }
        if input[py][px] == '+' {
            if direction % 2 == 0 {
                if input[py+1][px] != ' ' {
                    direction = 1
                } else {
                    direction = 3
                }
            } else {
                if input[py][px+1] != ' ' {
                    direction = 4
                } else {
                    direction = 2
                }
            }
        }
        if int(input[py][px]) > 64 && int(input[py][px]) < 91 {
            output += string(input[py][px])
        }
    }
    fmt.Println(output)
    fmt.Println(steps)
}
