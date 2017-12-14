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
    s := string(b)
    input := strings.Split(s, "\n")

    delay := 0
    caught := true
    for caught{
    	caught = false
    	position := 0
    	delay++
    	for _, text := range input {
	    	if text == "" {
	    		break
	    	}
	        wall := strings.Fields(strings.Replace(text, ":", "", -1))
	        v, _ := strconv.Atoi(wall[0])
	        v2, _ := strconv.Atoi(wall[1])
	        for position < v {
	            position++
	        }
	        scannerp := (position+delay)%((v2-1)*2)
	        if scannerp == 0 {
	            caught = true
	            break
	        }
	    }
    }
    fmt.Print(delay)
}