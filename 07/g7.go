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

    scanner := bufio.NewScanner(file)
    nodes := [][]string{}
    child := []string{}
    for scanner.Scan() {
        words := strings.Fields(strings.Replace(scanner.Text(), ",", "", -1))
        words[1] = strings.Trim(words[1], "()")
        if len(words) > 3 {
            child = append(child, words[3:]...)
            nodes = append(nodes, append(words[:2], words[3:]...))
        } else {
            nodes = append(nodes, words[:2])
        }
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
            if !strings.HasSuffix(e[1], "!") { //if not flagged
                //look for child weights
                ctotal := []int{}
                allflagged := true
                for _, f := range nodes {
                    //search all nodes
                    for _, g := range e[2:] {
                        if f[0] == g { //if child found
                            if strings.HasSuffix(f[1], "!") { //if child is already summed
                                v, _ := strconv.Atoi(strings.Trim(f[1], "!")) //append
                                if len(ctotal) > 1 {
                                    if v != ctotal[0] { //find imbalance
                                        fmt.Println(f)
                                    }
                                }
                                ctotal = append(ctotal, v)
                            } else {
                                allflagged = false
                            }
                        }
                    }
                }
                if allflagged { //all children were summed
                    v, _ := strconv.Atoi(e[1])
                    sum := 0
                    eq := ctotal[0]
                    for _, q := range ctotal {
                        if q != eq { 
                            fmt.Println(e)
                            fmt.Println(ctotal)
                            run = false
                        }
                        sum += q
                    }
                    e[1] = strconv.Itoa(v + sum) + "!" //sum
                    if e[0] == root { //if we hit root, stop
                        run = false
                    }
                }
            }   
        }
    }
    //fmt.Println(nodes)
}