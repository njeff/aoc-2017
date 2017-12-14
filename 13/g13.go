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

    position := 0
    output := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        wall := strings.Fields(strings.Replace(scanner.Text(), ":", "", -1))
        v, _ := strconv.Atoi(wall[0])
        v2, _ := strconv.Atoi(wall[1])
        for position < v {
            position++
        }
        scannerp := position%((v2-1)*2)
        if scannerp == 0 {
            output += v * v2
        }
    }

    fmt.Println(output)
}