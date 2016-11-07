package main

import "fmt"

// Closing the channel will signal that there will be no more values sending to channels
func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    go func() {
        for {
            j, more := <-jobs // In this special 2-value form of receive, 
                              //the more value will be false if jobs has been closed and all values in the channel have already been received.
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("send job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")
    <-done
}
