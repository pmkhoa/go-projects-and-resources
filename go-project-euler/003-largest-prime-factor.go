package main

import "fmt"

func largestPrimeFactor(number int) int {
    p := 0
    for i := 3; i*i < number; i += 2 {
        if number % i == 0 {
            p = i
            for number % i == 0 {
                number /= i
            }
        }
    }
    if number > 1 {
        p = number
    }
    return p
}

func main() {
    fmt.Println(largestPrimeFactor(600851475143))
}
