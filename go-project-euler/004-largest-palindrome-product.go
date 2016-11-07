package main

import "fmt"

func palindromeFrom3Digits() []int {
    prod := 0
    var palindromeList []int
    for i := 99999; i > 90000; i-- {
        for j := 99999; j > 90000; j -- {
            prod = i * j
            if isPalindrome( prod ) {
                palindromeList = append(palindromeList, prod)
            }
        }
    }
    return palindromeList
}

func getMax(a []int) int {
    var max = 0
    for i := 0; i < len(a) - 1; i++ {
        if ( a[i] > max ) {
            max = a[i]
        }
    }
    return max
}

func isPalindrome(number int) bool {
    reverse := 0
    temp := number
    remainder := 0
    for temp != 0 {
        remainder = temp % 10
        reverse = reverse * 10 + remainder
        temp /= 10
    }
    if reverse == number {
        return true
    }
    return false
}

func main() {
    max := getMax(palindromeFrom3Digits())
    fmt.Println(max)
}
