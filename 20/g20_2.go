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
    
    for steps := 0; steps < 1E3; steps++ {
        to_remove := make(map[string]int)
        idx_to_r := []int{}
        removed := 0
        for i := 0; i<len(particles); i++{
            check := false
            for j := 0; j<9; j++ {
                if particles[i][j] != 0 {
                    check = true
                    break
                }
            }

            if check {
                p := strconv.Itoa(particles[i][0]) + strconv.Itoa(particles[i][1]) + strconv.Itoa(particles[i][2])
                if to_remove[p] == 0 {
                    to_remove[p] = i+1 //case of index 0
                } else {
                    idx_to_r = append(idx_to_r, i)
                    idx_to_r = append(idx_to_r, to_remove[p] - 1)
                }

                for j := 0; j<3; j++ { //update
                    particles[i][j+3] += particles[i][j+6]
                    particles[i][j] += particles[i][j+3]
                }
            } else {
                removed++
            }
        }

        for _, e:= range idx_to_r {
            for j := 0; j<9; j++ {
                particles[e][j] = 0
            }
        }
        fmt.Println(len(particles)-removed)
    }
}

func remove(arr [][]int, i int) [][]int {
    arr[i] = arr[0]
    return arr[1:]
}
