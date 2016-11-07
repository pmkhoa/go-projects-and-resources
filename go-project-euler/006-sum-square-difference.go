package main

import "fmt"

func getSumOfSquares(upperNumber int) int {
    sum := 0
    for i := 1; i <= upperNumber; i++ {
        sum += i*i
    }
    return sum
}

func getSquareOfSum(upperNumber int) int {
    sum := (upperNumber * (upperNumber+1))/2
    return sum*sum
}

func diff(squareOfSum int, sumOfSquare int) int {
    return squareOfSum - sumOfSquare
}

func main(){
    var sumOfSquare = getSumOfSquares(100)
    var squareOfSum = getSquareOfSum(100)
    fmt.Println(squareOfSum)
    fmt.Println(diff(squareOfSum, sumOfSquare))
}
