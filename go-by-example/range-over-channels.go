package main

import "fmt"

func main() {
    queue := make(chan string, 3)
    queue <- "one"
    queue <- "two"
    close(queue) // if not closing the queue, we will be block by trying to receive no value from channel

    // iterate over values in the queue channel
    for elem := range queue {
        fmt.Println(elem)
    }
}
