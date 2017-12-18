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

    reg := make([]map[string]int, 2)
    reg[0] = make(map[string]int)
    reg[1] = make(map[string]int)
    queue := make([][]int, 2)
    queue[0] = make([]int, 0)
    queue[1] = make([]int, 0)

    reg[1]["p"] = 1
    pc := []int{0,0}
    p1s := 0
    run := true

    for run && (pc[1] < len(input) || pc[0] < len(input)){
        for i := 0; i<2; i++ {
            if pc[i] >= len(input) {
                continue
            }
            line := input[pc[i]]
            v := strings.Split(line, " ")

            n1, e1 := strconv.Atoi(v[1])
            n2 := 0
            if e1 != nil {
                n1 = reg[i][v[1]]
            }

            if len(v) > 2 {
                n2t, e2 := strconv.Atoi(v[2])
                if e2 != nil {
                    n2 = reg[i][v[2]]
                } else {
                    n2 = n2t
                }
            }

            switch v[0] {
            case "set":
                reg[i][v[1]] = n2
            case "add":
                reg[i][v[1]] += n2
            case "mul":
                reg[i][v[1]] *= n2
            case "mod":
                reg[i][v[1]] %= n2
            case "snd":
                queue[1-i] = append(queue[1-i], n1)
                if i == 1 {
                    p1s++
                }
            case "rcv":
                if len(queue[i]) != 0 {
                    reg[i][v[1]] = queue[i][0]
                    queue[i] = queue[i][1:]
                } else {
                    pc[i]--
                    fmt.Println(p1s)
                }
            case "jgz":
                if n1 > 0 {
                    pc[i] += n2 - 1
                }
            }
            pc[i]++
        }
    }
}
