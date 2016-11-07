package main

import "fmt"
import "math"

// We get the list of primes under specific number
func getPrimesUnder(number int) []int {
    var results []int
    counter := 0
    for i := 2; i <= number; i++ {
        for j := 1; j <= int(math.Sqrt(float64(i))); j++ {
            if i % j == 0 {
                counter ++
            }
            if counter >= 2 {
                break
            }
        }
        if ( counter < 2 ) {
            results = append(results, i)
        }
        counter = 0
    }
    return results
}

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

func sumPrimes(primes []int) int {
    sum := 0
    for _, value := range primes {
        sum = sum + value
    }
    return sum
}

func main() {
    var primes = getPrimesUnder(2000000)
    fmt.Println(primes)
    fmt.Println(sumPrimes(primes))
}
