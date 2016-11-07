package main

import "fmt"

func get_C() int {
    a, b := get_AB()
    var c = 1000 - a - b
    return c
}

func get_AB() (int, int) {
    for a := 1; a < 1000; a++ {
        for b := 1; b < 1000; b++ {
            // Do little calculation to substitute c. We got this equation.
            var cond = 2000*a + 2000*b - 2*a*b - 1000*1000
            if cond == 0 {
                return a, b
            }
            cond = 0
        }
    }
    return 0,0
}

func main() {
    a, b := get_AB()
    c := get_C()
    fmt.Println(a*b*c)
}
