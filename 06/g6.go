package main

import (
    "fmt"
    "strings"
    "strconv"
)

func main() {
    input := "5 1 10 0 1 7 13 14 3 12 8 10 7 12 0 6"
    split := strings.Fields(input)

    banks := []int{}
    for _, e := range split {
        j, err := strconv.Atoi(e)
        if err != nil {
            panic(err)
        }
        banks = append(banks, j)
    }
    //fmt.Println(banks)
    cycles := 0
    prev := [][]int{}

    cont := true
    for cont {
        max := 0
        maxi := 0
        //get max
        for i, e := range banks {
            if e > max {
                max = e
                maxi = i
            }
        }
        banks[maxi] = 0
        index := (maxi+1)%len(banks)
        for i := 0; i< max; i++ {
            banks[index]++;
            index = (index+1)%len(banks)
        }

        //fmt.Println(banks)
        cycles++;
        for i, e := range prev {
            flag := true
            for j, v := range e {
                if v != banks[j] {
                    flag = false
                }
            }
            if flag {
                cont = false
                fmt.Println(cycles - i - 1)
            }
        }
        temp := make([]int, len(banks))
        copy(temp, banks)
        prev = append(prev,temp)
    }

    fmt.Println(cycles)
}