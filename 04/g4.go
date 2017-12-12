package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "sort"
)

func sortString(s string) string {
    split := strings.Split(s, "")
    sort.Strings(split)
    return strings.Join(split, "")
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Print(err)
    }
    defer file.Close()

    valid := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        words := strings.Split(scanner.Text(), " ")
        // for i, _ := range words {
        //     words[i] = sortString(words[i])
        // }
        allowed := true
        for i, e := range words {
            for _, f := range words[i+1:] {
                if e == f {
                    allowed = false
                }
            }
        }
        if allowed {
            valid++
        }
    }

    fmt.Print(valid)
}