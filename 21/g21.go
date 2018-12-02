package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "math"
)

func main() {
    b, err := ioutil.ReadFile("input.txt")
    if err != nil {
        fmt.Print(err)
    }
    s := strings.Replace(strings.TrimSpace(string(b)), "/", "", -1)
    input := strings.Split(s, "\n")

    enhance := make(map[string]string)
    for _, e := range input {
        ele := strings.Split(e, " => ")
        enhance[ele[0]] = ele[1]
        for i := 0; i<4; i++ {
            for j := 0; j<3; j++ {
                if len(ele[0]) == 4 {
                    t := f2(rot2(ele[0],i),j)
                    enhance[t] = ele[1]
                } else {
                    t := f3(rot3(ele[0],i),j)
                    enhance[t] = ele[1]
                }
            }
        }
    }

    p := ".#...####"

    for i := 0; i < 18; i++ {
        p = en(p,enhance)
        fmt.Println(strings.Count(p,"#"))
    }
}

func en(s string, en map[string]string) string {
    output := ""
    el := int(math.Sqrt(float64(len(s)))) //edge length
    if el%2 == 0 {
        //do 2x2
        for i := 0; i<el; i+=2 {
            temp := []string{"","",""}
            for j := 0; j<el; j+=2 { //build row
                //build 2x2
                t := string(s[i*el+j:i*el+j+2]) + string(s[(i+1)*el+j:(i+1)*el+j+2])
                temp[0] += string(en[t][:3])
                temp[1] += string(en[t][3:6])
                temp[2] += string(en[t][6:9])
            }
            output += temp[0] + temp[1] + temp[2]
        }
    } else {
        //do 3x3
        for i := 0; i<el; i+=3 {
            temp := []string{"","","",""}
            for j := 0; j<el; j+=3 {
                //build 3x3
                t := string(s[i*el+j:i*el+j+3]) + string(s[(i+1)*el+j:(i+1)*el+j+3]) + string(s[(i+2)*el+j:(i+2)*el+j+3])
                temp[0] += string(en[t][:4]) //generate rows
                temp[1] += string(en[t][4:8])
                temp[2] += string(en[t][8:12])
                temp[3] += string(en[t][12:16])
            }
            output += temp[0] + temp[1] + temp[2] + temp[3] //append 4 at a time
        }
    }
    return output
}

func f2(s string, a int) string {
    if a == 0 {
        return s
    }
    o := []rune(s)
    if a == 1 {
        o[0], o[1], o[2], o[3] = o[1], o[0], o[3], o[2]
    } else {
        o[0], o[1], o[2], o[3] = o[2], o[3], o[0], o[1]
    }
    return string(o)
}

func f3(s string, a int) string {
    if a == 0 {
        return s
    }
    o := []rune(s)
    if a == 1 {
        o[0], o[3], o[6], o[2], o[5], o[8] = o[2], o[5], o[8], o[0], o[3], o[6]
    } else {
        o[0], o[1], o[2], o[6], o[7], o[8] = o[6], o[7], o[8], o[0], o[1], o[2]
    }
    return string(o)
}

func rot2(s string, n int) string {
    o := []rune(s)
    for i := 0; i<n; i++ {
        o[0], o[1], o[2], o[3] = o[1], o[3], o[0], o[2]
    }
    return string(o)
}

func rot3(s string, n int) string {
    o := []rune(s)
    for i := 0; i<n; i++ {
        o[0], o[1], o[2], o[3], o[5], o[6], o[7], o[8] = o[2], o[5], o[8], o[1], o[7], o[0], o[3], o[6]
    }
    return string(o)
}
