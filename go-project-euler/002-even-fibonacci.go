package main

import "fmt"

func fibonacci() func() int {
    a, b := 0, 1
    return func() int {
        temp := a + b
        b = a
        a = temp
        return a
    }
}

func main() {
    f := fibonacci()
    sum := 0
    curr := 0
    for i := 0; sum < 4000000; i++ {
        curr += f() 
        if curr % 2 == 0 {
            sum += curr
        }
    }
    fmt.Println(sum)
}
