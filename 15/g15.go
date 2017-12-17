package main

import (
    "fmt"
)

func main() {
    gena := 289
    genb := 629

    count := 0
    va := gena
    vb := genb
    for i := 0; i<=40000000; i++ {
        va = (va * 16807)%2147483647
        vb = (vb * 48271)%2147483647
        if va%65536 == vb%65536 {
            count++
        }
   }
    fmt.Println(count)

    va = gena
    vb = genb
    count = 0
    for i := 0; i<=5000000; i++ {
        va = (va * 16807)%2147483647
        vb = (vb * 48271)%2147483647

        for va%4!=0 {
            va = (va * 16807)%2147483647
        }
        for vb%8!=0 {
            vb = (vb * 48271)%2147483647
        }
        if va%65536 == vb%65536 {
            count++
        }
   }
    fmt.Println(count)
}
