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

    value := 0
    total := 0
    trash := false
    skip := false
    trashcount := 0
    
    for _, e := range s {
        if skip {
            skip = false
        } else {
            c := string(e)
            switch c {
            case "{":
                if !trash {
                    value++
                }
            case "<":
                trash = true
            case ">":
                trash = false
                trashcount--
            case "!":
                skip = true
            case "}":
                if !trash {
                    value--
                }
            }
            if trash && !skip {
                trashcount++;
            }
            if !trash && c == "{" {
                total += value
            }
        }
    }
    fmt.Println(total)
    fmt.Println(trashcount)
}