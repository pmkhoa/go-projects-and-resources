package main

import (
    "fmt"
)

func main() {
    messages := make(chan string)
    signals := make(chan bool)

    select {
    case msg := <-messages:
        fmt.Println("Message received: ", msg)
    default:
        fmt.Println("No message received")
    }

    msg := "hi"
    select {
    case messages <- msg:
        fmt.Println("send message: ", msg)
    default:
        fmt.Println("no message sent")
    }

    select {
    case msg := <-messages:
        fmt.Println("Message received ", msg)
    case sig := <-signals:
        fmt.Println("Received signal ", sig)
    default:
        fmt.Println("no activity")
    }
}
