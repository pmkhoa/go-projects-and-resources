package main

import "fmt"
import "math/big"

func main() {
    sum := 0
    temp := new(big.Int)
    pow := new(big.Int)
    pow.SetString("1", 10)

    for i := 0; i < 1000; i ++ {
        temp.SetString("2", 10)
        pow = pow.Mul(pow, temp)
        // fmt.Println(pow)
    }
    for _, val := range pow.String() {
        sum += int(val - '0')
    }
    fmt.Println(sum)
}
