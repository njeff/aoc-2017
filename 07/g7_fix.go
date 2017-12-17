package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

type node struct {
    value int
    children []string
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Print(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    nodes := make(map[string]node)
    for scanner.Scan() {
        words := strings.Fields(strings.Replace(scanner.Text(), ",", "", -1))
        words[1] = strings.Trim(words[1], "()")
    }

    //find the root
    root := ""
    for _, e := range nodes {
        found := false
        for _, f := range child {
            if f == e[0] {
                found = true
                break
            }
        }
        if !found {
            root = e[0]
            fmt.Println("Root: " + e[0])
            break
        }
    }

    //flag all leaves
    for i, e := range nodes {
        if len(e) == 2 {
            nodes[i][1] = nodes[i][1] + "!"
        }
    }

    run := true
    for run {
        for _, e := range nodes {
            
        }
    }[]
}
