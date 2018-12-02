package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
    "math"
)

func main() {
    b, err := ioutil.ReadFile("input.txt")
    if err != nil {
        fmt.Print(err)
    }
    s := strings.TrimSpace(string(b))
    input := strings.Split(s, "\n")

    particles := [][]int{} //every particle in this format:
    //px, py, pz, vx, vy, vz, ax, ay, az

    for _, line := range input {
        v := strings.Split(line, ", ")
        p := []int{}
        for _, e:= range v {
            vals := strings.Split(strings.Trim(e[3:], ">"), ",")
            for _, f := range vals {
                n, _ := strconv.Atoi(f)
                p = append(p, n)
            }
        }
        particles = append(particles,p)
    }
    
    last_c := 0
    closest := 0
    steady := 0

    for true {
        lowest_dist := 1E10
        for i := 0; i<len(particles); i++ {
            distance := math.Abs(float64(particles[i][0])) + math.Abs(float64(particles[i][1])) + math.Abs(float64(particles[i][2]))
            if distance < lowest_dist {
                closest = i
                lowest_dist = distance
            }
            for j := 0; j<3; j++ { //update
                particles[i][j+3] += particles[i][j+6]
                particles[i][j] += particles[i][j+3]
            }
        }
        if last_c == closest {
            steady++
        } else {
            steady = 0
        }
        if steady > 1E4 {
            break
        }
        last_c = closest
    }
    fmt.Println(closest)
}

func remove(arr [][]int, i int) [][]int {
    arr[i] = arr[0]
    return arr[1:]
}
