package main

import (
    "fmt"
    "strings"
    "strconv"
)

func main() {
    input := "225,171,131,2,35,5,0,13,1,246,54,97,255,98,254,110"
    l := strings.Split(input,",")

    list := []int{}

    for i := 0; i<256; i++ {
        list = append(list, i)
    }
    position := 0
    skip := 0

    for _, e := range l {
        v, _ := strconv.Atoi(e)
        reverse(list, position, position + v-1)
        position += (v + skip)%256
        skip += 1
    }
    fmt.Println(list[0] * list[1])
}

func reverse(lst []int, start int, end int) {
    i := start%len(lst)
    j := end%len(lst)
    for k := 0; k<(end-start+1)/2; k++ {
        lst[i], lst[j] = lst[j], lst[i]
        i = (i+1)%len(lst)
        j--
        if j < 0 {
            j = len(lst)-1
        }
    }
}