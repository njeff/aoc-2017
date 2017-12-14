package main

import (
    "fmt"
    "strings"
    "strconv"
)

func main() {
    input := "wenycdww"
    output := ""

    for l := 0; l<128; l++ {
        input := input + "-" + strconv.Itoa(l)
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
            output += fmt.Sprintf("%08b", e)
        }
        output += "\n"
    }

    total := 0
    for _, e := range output {
        if e == '1' {
            total++
        }
    }

    so := strings.Split(output, "\n")
    sr := [][]int{}
    sr = append(sr, make([]int, 130))
    for _, e := range so {
        if e != "" {
            row := []int{}
            row = append(row, 0) //pad sides
            for _, b := range e {
                row = append(row, int(b - 48))
            }
            row = append(row, 0) //pad sides
            sr = append(sr, row)
        }
    }
    sr = append(sr, make([]int, 130))

    groups := 0
    for row := 1; row < len(sr)-1; row++ {
        for col := 1; col < len(sr[0])-1; col++ {
            if sr[row][col] == 1 {
                groups++
                remove(sr,row,col)
            }
        }
    }

    fmt.Println(total)
    fmt.Println(groups)
}

func remove(lst [][]int, x int, y int){
    if lst[x][y] == 0 {
        return
    }
    lst[x][y] = 0
    remove(lst, x-1, y)
    remove(lst, x+1, y)
    remove(lst, x, y-1)
    remove(lst, x, y+1)
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