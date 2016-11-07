package main

import "fmt"
import "math"

func isPrime(number int) bool {
    counter := 0
    for i := 1; i <= int(math.Sqrt(float64(number))); i++ {
        if number % i == 0 {
            counter ++
        }
        if counter >= 2 {
            return false
        }
    }
    return true 
}

func getPrimesAtPosition(upperNumber int) int {
    counter := 0
    i := 0
    for counter <= upperNumber {
        i++
        if isPrime(i) {
            counter++
        }
    }
    return i
}

func main() {
    fmt.Println(getPrimesAtPosition(100001))
}
