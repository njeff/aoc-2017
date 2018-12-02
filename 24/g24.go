package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

func main() {
    b, err := ioutil.ReadFile("input.txt")
    if err != nil {
        fmt.Print(err)
    }
    s := strings.TrimSpace(string(b))
    input := strings.Split(s, "\n")

    seg := [][]int{}
    for _, e := range input {
        ele := strings.Split(e, "/")
        v1, _ := strconv.Atoi(ele[0])
        v2, _ := strconv.Atoi(ele[1])
        seg = append(seg, []int{v1,v2, v1+v2})
    }
    used := make([]int, len(seg))

    fmt.Println(sel(seg,used,0,0))
    fmt.Println(ls)
}

var longest = 0
var ls = 0
func sel(seg [][]int, used []int, tail int, max int) int {
    e := 0
    for i := 0; i<len(used); i++ {
        e += used[i]
    }
    if e > longest {
        longest = e
        ls = max
    }
    if e == longest {
        if max > ls {
            ls = max
        }
    }
    tag := false
    v := make([]int, len(used))
    for i := 0; i<len(seg); i++ {
        if used[i] == 0 {
            if seg[i][0] == tail {
                tag = true
                used[i] = 1
                v[i] = sel(seg,used,seg[i][1],max+seg[i][2])
                used[i] = 0
            } else if seg[i][1] == tail {
                tag = true
                used[i] = 1
                v[i] = sel(seg,used,seg[i][0],max+seg[i][2])
                used[i] = 0
            }
        }
    }
    if !tag {
        return max
    } else {
        return maxs(v)
    }
}

func maxs(v []int) int {
    m := 0
    for _, e := range v {
        if e > m {
            m = e
        }
    }
    return m
}
