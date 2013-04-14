package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "math"
    "os"
    "strconv"
    "strings"
)

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
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

    var success []int
    var num_test_cases = 0

    for {
        if idx >= len(input) {
            break
        }
        if idx != 0 {

            items := strings.Split(input[idx], " ")

            if len(items) > 1 {
                start, _ := strconv.Atoi(items[0])
                end, _ := strconv.Atoi(items[1])

                for i := start; i <= end; i++ {
                    if strconv.Itoa(i) == Reverse(strconv.Itoa(i)) {
                        f := float64(i)
                        whole, frac := math.Modf(math.Sqrt(f))
                        if frac == 0 {
                            if strconv.Itoa(int(whole)) == Reverse(strconv.Itoa(int(whole))) {
                                success[idx-1]++
                            }
                        }
                    }
                }
            }

        } else {
            num_test_cases, _ = strconv.Atoi(input[idx])
            success = make([]int, num_test_cases, num_test_cases)
        }

        idx++
    }

    for i := 0; i < num_test_cases; i++ {
        fmt.Println("Case #" + strconv.Itoa(i+1) + ": " + strconv.Itoa(success[i]))
    }
}
