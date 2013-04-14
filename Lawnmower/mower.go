package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
    "strings"
)

type Lawn struct {
    Pattern [][]int
}

func ShrinkSlice(pattern [][]int) [][]int {
    total := len(pattern)
    for i := 0; i < total; i++ {
        if len(pattern[i]) == 0 {
            copy(pattern[i:], pattern[i+1:])
            pattern = pattern[:len(pattern)-1]
        }
    }
    return pattern
}

func (l Lawn) Solve() string {
    fmt.Println("solve")
    idx := 0

    for {
        if idx >= len(l.Pattern) {
            break
        }

        var j = l.Pattern[idx][0] * len(l.Pattern[idx])
        var total = 0
        for _, inner := range l.Pattern[idx] {
            total += inner
        }

        if total == j {
            l.Pattern[idx] = nil
        }
        idx++
    }

    l.Pattern = ShrinkSlice(l.Pattern)

    if len(l.Pattern) == 0 {
        return "YES"
    }

    fmt.Println(l.Pattern)
    var rows = len(l.Pattern)
    var columns = len(l.Pattern[0])

    fmt.Println("columns", columns, "rows", rows)
    //fmt.Println("Printing the first column")
    for j := 0; j < columns; j++ {
        total := 0
        sum := 0
        for i := 0; i < rows; i++ {
            if i == 0 {
                total = l.Pattern[i][j] * rows
            }
            //fmt.Printf("%d", l.Pattern[i][j])
            sum += l.Pattern[i][j]
        }
        if sum != total {
            return "NO"
        }
        //fmt.Printf(" = %d (%d)", sum, total)

    }

    return "YES"
}

func main() {
    args := os.Args
    filename := args[1]

    b, err := ioutil.ReadFile(filename)
    if err != nil {
        panic(err)
    }

    idx := 0
    contents := bytes.NewBuffer(b)
    input := strings.Split(contents.String(), "\n")

    var num_test_cases = 0
    lawns := []Lawn{}

    for {
        if idx >= len(input) {
            break
        }
        if idx != 0 {
            size := strings.Split(input[idx], " ")

            if len(size) > 1 {

                fmt.Println("size =", size)

                height, _ := strconv.Atoi(size[0])
                width, _ := strconv.Atoi(size[1])

                lawn := Lawn{Pattern: make([][]int, height)}

                for i := 1; i <= height; i++ {

                    line := strings.Split(input[idx+i], " ")
                    lawn.Pattern[i-1] = make([]int, width)

                    for j := 0; j < width; j++ {
                        value, _ := strconv.Atoi(line[j])
                        lawn.Pattern[i-1][j] = value
                    }
                }

                idx = idx + height

                lawns = append(lawns, lawn)

            }

        } else {
            num_test_cases, _ = strconv.Atoi(input[idx])
        }

        idx++
    }

    for i := 0; i < num_test_cases; i++ {
        fmt.Println("Case #" + strconv.Itoa(i+1) + ":" + lawns[i].Solve())
    }
}
