package main

import (
    "fmt"
)

func main() {
    input := "225,171,131,2,35,5,0,13,1,246,54,97,255,98,254,110"

    list := []int{}

    for i := 0; i<256; i++ {
        list = append(list, i)
    }
    position := 0
    skip := 0

    for i := 0; i < 64; i++ {
        for _, e := range input {
            v := int(e)
            reverse(list, position, position + v-1)
            position += (v + skip)%256
            skip += 1
        }
        for _, e := range []int{17, 31, 73, 47, 23} {
            v := e
            reverse(list, position, position + v-1)
            position += (v + skip)%256
            skip += 1
        }
    }

    dense := []int{}
    for i := 0; i<16; i++ {
        v := list[i*16]
        for j:= 1; j<16; j++ {
            v ^= list[i*16+j]
        }
        dense = append(dense, v)
    }

    for _, e := range dense {
        fmt.Print(fmt.Sprintf("%02x", e))
    }
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