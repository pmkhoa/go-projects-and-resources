package main

import "fmt"
import "math/rand"
import "time"

func pinger(c chan string) {
	c <- "ping"
}

func ponger(c1 chan string, c2 chan string) {
	msg := <-c1
	fmt.Println(msg)
	c2 <- "pong"
}

func printer(c chan string) {
	fmt.Println(<-c)
}

func writer(msg string, c chan string) {
	for {
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		c <- msg
	}
}

func main() {
	c1, c2 := make(chan string), make(chan string)
	// Using go routine make sure all of three tasks
	// start running at the same time
	// However, as the channel requires that ping comes to channel first,
	// Then it comes out, next to pong comes in channel
	// after that it comes out.
	// go pinger(c1)
	// go ponger(c1, c2)
	// go printer(c2)

	go writer("Writer 1", c1)
	go writer("Writer 2", c2)

	for {
		select {
		case msg1 := <-c1:
			fmt.Println("Message 1", msg1)
		case msg2 := <-c2:
			fmt.Println("Message 2", msg2)
		case <-time.After(time.Second):
			fmt.Println("Timeout")
		}
	}

	var input string
	fmt.Scanln(&input)
}
