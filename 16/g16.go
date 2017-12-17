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
    s := string(b[0:len(b)-1])
    input := strings.Split(s, ",")

    prog := []rune("abcdefghijklmnop")

    memo := []string{}

    for i := 0; i<1000000000; i++ {
        memo = append(memo, string(prog))
        for _, e := range input {
            switch e[0] {
            case 's':
                v, _ := strconv.Atoi(e[1:])
                spin(prog,v)
            case 'x':
                positions := strings.Split(e[1:], "/")
                p1, _ := strconv.Atoi(positions[0])
                p2, _ := strconv.Atoi(positions[1])
                prog[p1], prog[p2] = prog[p2], prog[p1]
            case 'p':
                a := strings.Index(string(prog), string(e[1]))
                b := strings.Index(string(prog), string(e[3]))
                prog[a], prog[b] = prog[b], prog[a]
            }
        }
        if in(memo, string(prog)) {
            fmt.Println(memo[1])
            fmt.Println(string(memo[1E9%len(memo)]))
            break
        }
    }

}

func in(arr []string, s string) bool {
    for _, e := range arr {
        if e == s {
            return true
        }
    }
    return false
}

func spin(prog []rune, v int) {
    for i := 0; i<v; i++ {
        temp := prog[len(prog)-1]
        for j := len(prog)-1; j>0; j-- {
            prog[j] = prog[j-1]
        }
        prog[0] = temp
    }
}
