package main

import (
    "fmt"
    "math"
)

func main() {
    var input float64 = 289326
    ring := (math.Floor(math.Sqrt(float64(input-1)))+1)/2;
    lowerbound := (ring*2-1)*(ring*2-1);
    upperbound := (ring*2+1)*(ring*2+1);
    
    position := math.Mod((input - lowerbound)/(upperbound - lowerbound),0.25)/0.25;
    rd := ring*math.Abs(0.5-position)/0.5;
    fmt.Print(int(ring+rd+0.5)) //round up
}