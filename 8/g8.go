package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Print(err)
    }
    defer file.Close()

    m := make(map[string]int)

    scanner := bufio.NewScanner(file)
    max := 0
    for scanner.Scan() {
        inst := strings.Fields(scanner.Text())
        value, _ := strconv.Atoi(inst[2])
        if inst[1] != "inc" {
            value *= -1
        }
        flag := false
        val := m[inst[4]]
        comp, _ := strconv.Atoi(inst[6])
        switch inst[5] {
        case "!=":
            flag = (val != comp)
        case "<=":
            flag = (val <= comp)
        case ">=":
            flag = (val >= comp)
        case "==":
            flag = (val == comp)
        case "<":
            flag = (val < comp)
        case ">":
            flag = (val > comp)
        }
        if flag {
            m[inst[0]] += value
        }
        for _, v := range m {
            if v > max {
                max = v
            }
        }
    }
    fm := 0
    for _, v := range m {
        if v > fm {
            fm = v
        }
    }
    fmt.Println(fm)
    fmt.Println(max)
}