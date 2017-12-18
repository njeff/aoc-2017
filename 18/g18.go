package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

func main() {
    b, err := ioutil.ReadFile("input.txt")
    if err != nil {
        fmt.Print(err)
    }
    s := strings.TrimSpace(string(b))
    input := strings.Split(s, "\n")

    reg := make(map[string]int)
    pc := 0
    play := 0
    run := true

    for run && pc < len(input){
        line := input[pc]
        v := strings.Split(line, " ")

        n1, e1 := strconv.Atoi(v[1])
        n2 := 0
        if e1 != nil {
            n1 = reg[v[1]]
        }

        if len(v) > 2 {
            n2t, e2 := strconv.Atoi(v[2])
            if e2 != nil {
                n2 = reg[v[2]]
            } else {
                n2 = n2t
            }
        }

        switch v[0] {
        case "set":
            reg[v[1]] = n2
        case "add":
            reg[v[1]] += n2
        case "mul":
            reg[v[1]] *= n2
        case "mod":
            reg[v[1]] %= n2
        case "snd":
            play = n1
        case "rcv":
            if n1 != 0 {
                fmt.Println(play)
                run = false
            }
        case "jgz":
            if n1 > 0 {
                pc += n2 - 1
            }
        }
        pc++
    }
}

