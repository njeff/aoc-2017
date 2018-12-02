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

    tape := make(map[int]string)
    position := 0
    states := make(map[string][]string)

    state := strings.Trim(input[0],"Begin in state.")
    steps, _ := strconv.Atoi(strings.Trim(input[1],"Perform a diagnostic checksum after steps."))

    for i := 3; i<len(input); i+=10 {
        s := strings.Trim(input[i],"In state :")
        temp := make([]string, 6)
        for j := 0; j<2; j++ {
            temp[j*3] = strings.Trim(input[i+2+j*4], " - Write the value .")
            temp[j*3+1] = strings.Trim(strings.TrimPrefix(input[i+3+j*4], "    - Move one slot to the "),".")
            temp[j*3+2] = strings.Trim(strings.TrimPrefix(input[i+4+j*4], "    - Continue with state "),".")
        }
        fmt.Println(temp)
        states[s] = temp
    }

    for i := 0; i<steps; i++ {
        val := 0
        if tape[position] == "1" {
            val = 1
        }
        tape[position] = states[state][val*3]
        mv := -1
        if states[state][val*3+1] == "right" {
            mv = 1
        }
        position += mv
        state = states[state][val*3+2]
    }

    total := 0
    for _, v := range tape {
        if v == "1" {
            total++
        }
    }
    fmt.Println(total)
}

