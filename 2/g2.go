package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func checkline(s string) int {
    var max, min int = -1E7, 1E7
    words := strings.Fields(s)
    for _, e := range words {
        i, _ := strconv.Atoi(e)
        if i < min {
            min = i
        }
        if i > max {
            max = i
        }
    }
    return max-min
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Print(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    checksum := 0
    for scanner.Scan() {
        checksum += checkline(scanner.Text())
    }

    fmt.Print(checksum)
}