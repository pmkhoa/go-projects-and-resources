package main

import "fmt"

func main() {
    messages := make(chan string, 2)
    messages <- "Buffered"
    messages <- "Channel"

    // All goroutines are asleep. deadlock
    // messages <- "Channel Overflowed"

    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
