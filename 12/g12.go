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

    program := 0
    program_map := make(map[int][]int)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        connections := strings.Fields(strings.Replace(scanner.Text(), ",", "", -1))[2:]
        coni := []int{}
        for i := 0; i<len(connections); i++ {
            v, _ := strconv.Atoi(connections[i])
            coni = append(coni, v)
        }
        program_map[program] = coni
        program++       
    }

    seen := make(map[int]int)
    total0 := countnodes(program_map, 0, seen)
    groups := 1

    for len(program_map) != len(seen) {
        for i := 0; i < len(program_map); i++ {
            if seen[i] == 0 {
                countnodes(program_map, i, seen)
                break
            }
        }
        groups++
    }

    fmt.Println(total0)
    fmt.Println(groups)
}

//counts nodes in a tree
func countnodes(m map[int][]int, root int, seen map[int]int) int {
    //for all branches
    total := 1
    seen[root] = 1
    for _, e := range m[root] {
        if seen[e] == 0 {
            total += countnodes(m, e, seen)
        }
    }
    return total
}