package main
import "fmt"


// We get the list of primes under specific number
func getPrimesUnder(number int) []int {
    var results []int
    counter := 0
    for i := 1; i <= number; i++ {
        for j := 1; j < i; j++ {
            if i % j == 0 {
                counter ++
            }
        }
        if ( counter < 2 ) {
            results = append(results, i)
        }
        counter = 0
    }
    return results
}

// Find the largest prime compose number in that prime list
// in which it still under the limit number
func getLargestPrimeComposeNumbers(limit int, list []int) []int {
    var results []int
    var temp = 1
    for i := 0; i < len(list); i++ {
        for temp <= limit && list[i] != 1 {
            temp = temp * list[i]
        }
        results = append(results, temp/list[i])
        temp = 1
    }
    return results
}

// We return the final number from the prime list
func getFinalNumber(list []int) int {
    var result = 1
    for i := 0; i < len(list); i++ {
        result *= list[i]
    }
    return result
}

func main() {
    var a = getPrimesUnder(20)
    var primeList = getLargestPrimeComposeNumbers(20, a)
    fmt.Println(getFinalNumber(primeList))
}
